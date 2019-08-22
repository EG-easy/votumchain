package votum

import (
	"github.com/EG-easy/votumchain/x/votum/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
)

const (
	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	QuerierRoute      = types.QuerierRoute
	DefaultParamspace = types.DefaultParamspace
	DefaultCodespace  = types.DefaultCodespace

	StatusNil = gov.StatusNil
)

var (
	// TODO: Msgをここで定める
	ModuleCdc     = types.ModuleCdc
	RegisterCodec = types.RegisterCodec
)

type (
	MsgIssueToken     = types.MsgIssueToken
	MsgSubmitProposal = types.MsgSubmitProposal
	MsgDeposit        = types.MsgDeposit
	MsgVote           = types.MsgVote
)

//gov module functions
var (
	ParamKeyTable                 = gov.ParamKeyTable
	ParamStoreKeyVotingParams     = gov.ParamStoreKeyVotingParams
	ProposalsKeyPrefix            = gov.ProposalsKeyPrefix
	SplitActiveProposalQueueKey   = gov.SplitActiveProposalQueueKey
	SplitInactiveProposalQueueKey = gov.SplitInactiveProposalQueueKey
	DepositsKeyPrefix             = gov.DepositsKeyPrefix
	ErrUnknownProposal            = gov.ErrUnknownProposal
	ErrAlreadyActiveProposal      = gov.ErrAlreadyActiveProposal
	ErrAlreadyFinishedProposal    = gov.ErrAlreadyFinishedProposal
	ErrAddressNotStaked           = gov.ErrAddressNotStaked
	ErrInvalidProposalContent     = gov.ErrInvalidProposalContent
	ErrInvalidProposalType        = gov.ErrInvalidProposalType
	ErrInvalidVote                = gov.ErrInvalidVote
	ErrInvalidGenesis             = gov.ErrInvalidGenesis
	ErrNoProposalHandlerExists    = gov.ErrNoProposalHandlerExists
	StatusDepositPeriod           = gov.StatusDepositPeriod
	StatusVotingPeriod            = gov.StatusVotingPeriod
	StatusPassed                  = gov.StatusPassed
	StatusRejected                = gov.StatusRejected
	StatusFailed                  = gov.StatusFailed
	NewDeposit                    = gov.NewDeposit
	DepositsKey                   = gov.DepositsKey
	DepositKey                    = gov.DepositKey
	VotesKey                      = gov.VotesKey
	VoteKey                       = gov.VoteKey
	ValidVoteOption               = gov.ValidVoteOption
	NewVote                       = gov.NewVote
	NewProposal                   = gov.NewProposal
	ProposalKey                   = gov.ProposalKey
	ValidProposalStatus           = gov.ValidProposalStatus
	ProposalIDKey                 = gov.ProposalIDKey
	ParamStoreKeyDepositParams    = gov.ParamStoreKeyDepositParams
	ParamStoreKeyTallyParams      = gov.ParamStoreKeyTallyParams
	ActiveProposalQueueKey        = gov.ActiveProposalQueueKey
	InactiveProposalQueueKey      = gov.InactiveProposalQueueKey
	ActiveProposalByTimeKey       = gov.ActiveProposalByTimeKey
	InactiveProposalByTimeKey     = gov.InactiveProposalByTimeKey
	ErrInactiveProposal           = gov.ErrInactiveProposal

	VotesKeyPrefix              = gov.VotesKeyPrefix
	ActiveProposalQueuePrefix   = gov.ActiveProposalQueuePrefix
	InactiveProposalQueuePrefix = gov.InactiveProposalQueuePrefix
)

type (
	Router = gov.Router
	// MsgSubmitProposal = gov.MsgSubmitProposal
	// MsgDeposit        = gov.MsgDeposit
	// MsgVote           = gov.MsgVote
	DepositParams  = gov.DepositParams
	VotingParams   = gov.VotingParams
	TallyParams    = gov.TallyParams
	Proposal       = gov.Proposal
	Deposit        = gov.Deposit
	Vote           = gov.Vote
	Deposits       = gov.Deposits
	VoteOption     = gov.VoteOption
	Votes          = gov.Votes
	Content        = gov.Content
	Proposals      = gov.Proposals
	ProposalStatus = gov.ProposalStatus
)
