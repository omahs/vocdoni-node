package offchaindatahandler

import (
	"sync"

	"go.vocdoni.io/dvote/api/censusdb"
	"go.vocdoni.io/dvote/data/downloader"
	"go.vocdoni.io/dvote/log"
	"go.vocdoni.io/dvote/util"
	"go.vocdoni.io/dvote/vochain"
	"go.vocdoni.io/dvote/vochain/state"
	"go.vocdoni.io/dvote/vochain/transaction/vochaintx"
	"go.vocdoni.io/proto/build/go/models"
)

const (
	itemTypeExternalCensus = iota
	itemTypeRollingCensus
	itemTypeOrganizationMetadata
	itemTypeElectionMetadata
	itemTypeAccountMetadata
)

type importItem struct {
	itemType   int
	uri        string
	censusRoot string
	pid        []byte
}

// TBD: A startup process for importing on-going process census
// TBD: a mechanism for removing already finished census?

// OffChainDataHandler is a Vochain event handler aimed to fetch
// offchain data (usually on IPFS).
type OffChainDataHandler struct {
	vochain       *vochain.BaseApplication
	census        *censusdb.CensusDB
	storage       *downloader.Downloader
	queue         []importItem
	queueLock     sync.RWMutex
	importOnlyNew bool
	isFastSync    bool
}

// NewOffChainDataHandler creates a new instance of the off chain data downloader daemon.
// It will subscribe to Vochain events and perform data import.
func NewOffChainDataHandler(v *vochain.BaseApplication, d *downloader.Downloader,
	c *censusdb.CensusDB, importOnlyNew bool) *OffChainDataHandler {
	od := OffChainDataHandler{
		vochain:       v,
		census:        c,
		storage:       d,
		importOnlyNew: importOnlyNew,
		queue:         make([]importItem, 0),
	}
	v.State.AddEventListener(&od)
	return &od
}

func (c *OffChainDataHandler) Rollback() {
	c.queueLock.Lock()
	c.queue = make([]importItem, 0)
	c.isFastSync = c.vochain.IsSynchronizing()
	c.queueLock.Unlock()
}

// Commit is called when a new block is committed, so we execute the import actions
// enqueued by the event handlers (else the queues are reverted by calling Rollback).
func (c *OffChainDataHandler) Commit(height uint32) error {
	c.queueLock.Lock()
	defer c.queueLock.Unlock()
	for _, item := range c.queue {
		switch item.itemType {
		case itemTypeExternalCensus:
			log.Infof("importing external census %s", item.uri)
			// AddToQueue() writes to a channel that might be full, so we don't want to block the main thread.
			go c.enqueueOffchainCensus(item.censusRoot, item.uri)
		case itemTypeElectionMetadata, itemTypeAccountMetadata:
			log.Infof("importing metadata from %s", item.uri)
			go c.enqueueMetadata(item.uri)
		case itemTypeRollingCensus:
			log.Infof("importing rolling census for process %x", item.pid)
			c.importRollingCensus(item.pid)
		default:
			log.Errorf("unknown item %d", item.itemType)
		}
	}
	return nil
}

// OnProcess is triggered when a new election is created. It checks if the election contains offchain data
// that needs to be imported and enqueues it for being handled by Commit.
func (c *OffChainDataHandler) OnProcess(pid, eid []byte, censusRoot, censusURI string, txindex int32) {
	censusRoot = util.TrimHex(censusRoot)
	c.queueLock.Lock()
	defer c.queueLock.Unlock()
	if !c.importOnlyNew || !c.isFastSync {
		p, err := c.vochain.State.Process(pid, false)
		if err != nil || p == nil {
			log.Errorf("could get process from state: %v", err)
			return
		}
		// enqueue for import election metadata information
		if m := p.GetMetadata(); m != "" {
			log.Debugf("adding election metadata %s to queue", m)
			c.queue = append(c.queue, importItem{
				uri:      m,
				itemType: itemTypeElectionMetadata,
			})
		}
		// enqueue for download external census if needs to be imported
		if state.CensusOrigins[p.CensusOrigin].NeedsDownload && len(censusURI) > 0 {
			log.Infof("adding external census %s to queue", censusURI)
			c.queue = append(c.queue, importItem{
				censusRoot: censusRoot,
				uri:        censusURI,
				itemType:   itemTypeExternalCensus,
			})
		}
	}
}

// OnProcessStart is triggered when a process starts. It checks if the process contains a rolling census.
func (c *OffChainDataHandler) OnProcessesStart(pids [][]byte) {
	for _, pid := range pids {
		process, err := c.vochain.State.Process(pid, true)
		if err != nil {
			log.Errorf("could find process with pid %x: %v", pid, err)
			continue
		}
		// enqueue for import rolling census (zkSnarks voting with preregister enabled)
		if process.Mode.PreRegister && process.EnvelopeType.Anonymous {
			c.queueLock.Lock()
			log.Infof("adding rolling census for process %x to queue", pid)
			c.queue = append(c.queue, importItem{
				itemType: itemTypeRollingCensus,
				pid:      pid,
			})
			c.queueLock.Unlock()
		}
	}
}

// OnSetAccount is triggered when a new account is created or modifyied. If metadata info is present, it is enqueued.
func (c *OffChainDataHandler) OnSetAccount(addr []byte, account *state.Account) {
	c.queueLock.Lock()
	defer c.queueLock.Unlock()
	if !c.importOnlyNew || !c.isFastSync {
		// enqueue for import account metadata information
		if m := account.GetInfoURI(); m != "" {
			log.Debugf("adding account info metadata %s to queue", m)
			c.queue = append(c.queue, importItem{
				uri:      m,
				itemType: itemTypeAccountMetadata,
			})
		}
	}
}

// NOT USED but required for implementing the vochain.EventListener interface
func (c *OffChainDataHandler) OnCancel(pid []byte, txindex int32)                                 {}
func (c *OffChainDataHandler) OnVote(v *models.Vote, voterID state.VoterID, txindex int32)        {}
func (c *OffChainDataHandler) OnNewTx(tx *vochaintx.VochainTx, blockHeight uint32, txIndex int32) {}
func (c *OffChainDataHandler) OnProcessKeys(pid []byte, pub string, txindex int32)                {}
func (c *OffChainDataHandler) OnRevealKeys(pid []byte, priv string, txindex int32)                {}
func (c *OffChainDataHandler) OnProcessStatusChange(pid []byte, status models.ProcessStatus, txindex int32) {
}
func (c *OffChainDataHandler) OnTransferTokens(from, to []byte, amount uint64) {}
func (c *OffChainDataHandler) OnProcessResults(pid []byte, results *models.ProcessResult, txindex int32) {
}
