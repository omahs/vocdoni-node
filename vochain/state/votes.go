package state

import (
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/vocdoni/arbo"
	"go.vocdoni.io/dvote/crypto/ethereum"
	"go.vocdoni.io/dvote/db"
	"go.vocdoni.io/dvote/types"
	"go.vocdoni.io/proto/build/go/models"
	"google.golang.org/protobuf/proto"
)

var (
	// keys; not constants because of []byte
	voteCountKey = []byte("voteCount")
)

// VoteCount return the global vote count.
// When committed is false, the operation is executed also on not yet commited
// data from the currently open StateDB transaction.
// When committed is true, the operation is executed on the last commited version.
func (v *State) VoteCount(committed bool) (uint64, error) {
	if !committed {
		v.Tx.RLock()
		defer v.Tx.RUnlock()
	}
	noState := v.mainTreeViewer(committed).NoState()
	voteCountLE, err := noState.Get(voteCountKey)
	if errors.Is(err, db.ErrKeyNotFound) {
		return 0, nil
	} else if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint64(voteCountLE), nil
}

// voteCountInc increases by 1 the global vote count.
func (v *State) voteCountInc() error {
	noState := v.Tx.NoState()
	voteCountLE, err := noState.Get(voteCountKey)
	if errors.Is(err, db.ErrKeyNotFound) {
		voteCountLE = make([]byte, 8)
	} else if err != nil {
		return err
	}
	voteCount := binary.LittleEndian.Uint64(voteCountLE)
	voteCount++
	binary.LittleEndian.PutUint64(voteCountLE, voteCount)
	return noState.Set(voteCountKey, voteCountLE)
}

// AddVote adds a new vote to a process and call the even listeners to OnVote.
// If the vote already exists it will be overwritten and overwrite counter will be increased.
// Note that the vote is not committed to the StateDB until the StateDB transaction is committed.
// Note that the vote is not verified, so it is the caller responsibility to verify the vote.
func (v *State) AddVote(vote *models.Vote) error {
	vid, err := v.voteID(vote.ProcessId, vote.Nullifier)
	if err != nil {
		return err
	}
	// save block number
	vote.Height = v.CurrentHeight()
	voteBytes, err := proto.Marshal(vote)
	if err != nil {
		return fmt.Errorf("cannot marshal vote: %w", err)
	}

	sdbVote, err := v.Vote(vote.ProcessId, vote.Nullifier, false)
	if err != nil {
		if errors.Is(err, ErrVoteNotFound) {
			sdbVote = &models.StateDBVote{
				VoteHash:  ethereum.HashRaw(voteBytes),
				ProcessId: vote.ProcessId,
				Nullifier: vote.Nullifier,
			}
		} else {
			return err
		}
	} else {
		sdbVote.VoteHash = ethereum.HashRaw(voteBytes)
		if sdbVote.OverwriteCount != nil {
			*sdbVote.OverwriteCount++
		} else {
			sdbVote.OverwriteCount = new(uint32)
			*sdbVote.OverwriteCount = 1
		}
	}
	sdbVoteBytes, err := proto.Marshal(sdbVote)
	if err != nil {
		return fmt.Errorf("cannot marshal sdbVote: %w", err)
	}
	v.Tx.Lock()
	err = func() error {
		treeCfg := StateChildTreeCfg(ChildTreeVotes)
		if err := v.Tx.DeepSet(vid, sdbVoteBytes,
			StateTreeCfg(TreeProcess), treeCfg.WithKey(vote.ProcessId)); err != nil {
			return err
		}
		return v.voteCountInc()
	}()
	v.Tx.Unlock()
	if err != nil {
		return err
	}
	for _, l := range v.eventListeners {
		l.OnVote(vote, vote.VoterId, v.TxCounter())
	}
	return nil
}

// NOTE(Edu): Changed this from byte(processID+nullifier) to
// hash(processID+nullifier) to allow using it as a key in Arbo tree.
// voteID = hash(processID+nullifier)
func (v *State) voteID(pid, nullifier []byte) ([]byte, error) {
	if len(pid) != types.ProcessIDsize {
		return nil, fmt.Errorf("wrong processID size %d", len(pid))
	}
	if len(nullifier) != types.VoteNullifierSize {
		return nil, fmt.Errorf("wrong nullifier size %d", len(nullifier))
	}
	vid := sha256.New()
	vid.Write(pid)
	vid.Write(nullifier)
	return vid.Sum(nil), nil
}

// Vote returns the stored vote if exists. Returns ErrProcessNotFound if the
// process does not exist, ErrVoteNotFound if the vote does not exist.
// When committed is false, the operation is executed also on not yet commited
// data from the currently open StateDB transaction.
// When committed is true, the operation is executed on the last commited version.
func (v *State) Vote(processID, nullifier []byte, committed bool) (*models.StateDBVote, error) {
	vid, err := v.voteID(processID, nullifier)
	if err != nil {
		return nil, err
	}
	if !committed {
		// acquire a write lock, since DeepSubTree will create some temporary trees in memory
		// that might be read concurrently by DeliverTx path during block commit, leading to race #581
		// https://github.com/vocdoni/vocdoni-node/issues/581
		v.Tx.Lock()
		defer v.Tx.Unlock()
	}
	treeCfg := StateChildTreeCfg(ChildTreeVotes)
	votesTree, err := v.mainTreeViewer(committed).DeepSubTree(
		StateTreeCfg(TreeProcess), treeCfg.WithKey(processID))
	if errors.Is(err, arbo.ErrKeyNotFound) {
		return nil, ErrProcessNotFound
	} else if err != nil {
		return nil, err
	}
	sdbVoteBytes, err := votesTree.Get(vid)
	if errors.Is(err, arbo.ErrKeyNotFound) {
		return nil, ErrVoteNotFound
	} else if err != nil {
		return nil, err
	}
	var sdbVote models.StateDBVote
	if err := proto.Unmarshal(sdbVoteBytes, &sdbVote); err != nil {
		return nil, fmt.Errorf("cannot unmarshal sdbVote: %w", err)
	}
	return &sdbVote, nil
}

// VoteExists returns true if the envelope identified with voteID exists
// When committed is false, the operation is executed also on not yet commited
// data from the currently open StateDB transaction.
// When committed is true, the operation is executed on the last commited version.
func (v *State) VoteExists(processID, nullifier []byte, committed bool) (bool, error) {
	_, err := v.Vote(processID, nullifier, committed)
	if errors.Is(err, ErrProcessNotFound) {
		return false, nil
	} else if errors.Is(err, ErrVoteNotFound) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

// iterateVotes iterates fn over state tree entries with the processID prefix.
// When committed is false, the operation is executed also on not yet commited
// data from the currently open StateDB transaction.
// When committed is true, the operation is executed on the last commited version.
func (v *State) iterateVotes(processID []byte,
	fn func(vid []byte, sdbVote *models.StateDBVote) bool, committed bool) error {
	if !committed {
		v.Tx.RLock()
		defer v.Tx.RUnlock()
	}
	treeCfg := StateChildTreeCfg(ChildTreeVotes)
	votesTree, err := v.mainTreeViewer(committed).DeepSubTree(
		StateTreeCfg(TreeProcess), treeCfg.WithKey(processID))
	if err != nil {
		return err
	}
	var callbackErr error
	if err := votesTree.Iterate(func(key, value []byte) bool {
		var sdbVote models.StateDBVote
		if err := proto.Unmarshal(value, &sdbVote); err != nil {
			callbackErr = err
			return true
		}
		return fn(key, &sdbVote)
	}); err != nil {
		return err
	}
	if callbackErr != nil {
		return callbackErr
	}
	return nil
}

// CountVotes returns the number of votes registered for a given process id
// When committed is false, the operation is executed also on not yet commited
// data from the currently open StateDB transaction.
// When committed is true, the operation is executed on the last commited version.
func (v *State) CountVotes(processID []byte, committed bool) uint32 {
	var count uint32
	// TODO: Once statedb.TreeView.Size() works, replace this by that.
	v.iterateVotes(processID, func(vid []byte, sdbVote *models.StateDBVote) bool {
		count++
		return false
	}, committed)
	return count
}

// EnvelopeList returns a list of registered envelopes nullifiers given a processId
// When committed is false, the operation is executed also on not yet commited
// data from the currently open StateDB transaction.
// When committed is true, the operation is executed on the last commited version.
func (v *State) EnvelopeList(processID []byte, from, listSize int,
	committed bool) (nullifiers [][]byte) {
	idx := 0
	v.iterateVotes(processID, func(vid []byte, sdbVote *models.StateDBVote) bool {
		if idx >= from+listSize {
			return true
		}
		if idx >= from {
			nullifiers = append(nullifiers, sdbVote.Nullifier)
		}
		idx++
		return false
	}, committed)
	return nullifiers
}
