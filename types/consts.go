package types

import "time"

func Bool(b bool) *bool { return &b }

// These exported variables should be treated as constants, to be used in API
// responses which require *bool fields.
var (
	False = Bool(false)
	True  = Bool(true)
)

const (
	// All

	ProcessIDsize = 32
	// size of eth addr
	EntityIDsize = 20
	// legacy: in the past we used hash(addr)
	// this is a temporal work around to support both
	EntityIDsizeV2                 = 32
	VoteNullifierSize              = 32
	KeyIndexSeparator              = ":"
	EthereumConfirmationsThreshold = 6
	EntityResolverDomain           = "entity-resolver.vocdoni.eth"
	EntityMetaKey                  = "vnd.vocdoni.meta"
	EthereumReadTimeout            = 1 * time.Minute
	EthereumWriteTimeout           = 1 * time.Minute
	EthereumDialMaxRetry           = 10
	// Scrutinizer

	// ScrutinizerLiveProcessPrefix is used for sotring temporary results on live
	ScrutinizerLiveProcessPrefix = byte(0x21)
	// ScrutinizerEntityPrefix is the prefix for the storage entity keys
	ScrutinizerEntityPrefix = byte(0x22)
	// ScrutinizerResultsPrefix is the prefix of the storage results summary keys
	ScrutinizerResultsPrefix = byte(0x24)
	// ScrutinizerProcessEndingPrefix is the prefix for keep track of the processes ending on a specific block
	ScrutinizerProcessEndingPrefix = byte(0x25)

	// Vochain

	// PetitionSignStr contains the string that needs to match with the received vote type for petition-sign
	PetitionSignStr = "petition-sign"
	// PollVoteStr contains the string that needs to match with the received vote type for poll-vote
	PollVoteStr = "poll-vote"
	// EncryptedPollStr contains the string that needs to match with the received vote type for encrypted-poll
	EncryptedPollStr = "encrypted-poll"
	// SnarkVoteStr contains the string that needs to match with the received vote type for snark-vote
	SnarkVoteStr = "snark-vote"

	// List of transation names
	TxVote              = "vote"
	TxNewProcess        = "newProcess"
	TxCancelProcess     = "cancelProcess"
	TxAddValidator      = "addValidator"
	TxRemoveValidator   = "removeValidator"
	TxAddOracle         = "addOracle"
	TxRemoveOracle      = "removeOracle"
	TxAddProcessKeys    = "addProcessKeys"
	TxRevealProcessKeys = "revealProcessKeys"

	// MaxKeyIndex is the maxim number of allowed Encryption or Commitment keys
	MaxKeyIndex = 16
)

var (
	// PetitionSign contains the string that needs to match with the received vote type for petition-sign
	PetitionSign = [...]uint8{3, 0}
	// PollVote contains the string that needs to match with the received vote type for poll-vote
	PollVote = [...]uint8{3, 0}
	// EncryptedPoll contains the string that needs to match with the received vote type for encrypted-poll
	EncryptedPoll = [...]uint8{3, 4}
	// SnarkVote contains the string that needs to match with the received vote type for snark-vote
	SnarkVote = [...]uint8{3, 6}
)
