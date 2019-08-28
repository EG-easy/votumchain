package types

import "github.com/cosmos/cosmos-sdk/x/gov"

// query endpoints supported by the governance Querier
const (
	QueryParams    = gov.QueryParams
	QueryProposals = gov.QueryProposals
	QueryProposal  = gov.QueryProposal
	QueryDeposits  = gov.QueryDeposits
	QueryDeposit   = gov.QueryDeposit
	QueryVotes     = gov.QueryVotes
	QueryVote      = gov.QueryVote
	QueryTally     = gov.QueryTally

	ParamDeposit  = gov.ParamDeposit
	ParamVoting   = gov.ParamVoting
	ParamTallying = gov.ParamTallying
)

type (
	QueryProposalParams  = gov.QueryProposalParams
	QueryDepositParams   = gov.QueryDepositParams
	QueryVoteParams      = gov.QueryVoteParams
	QueryProposalsParams = gov.QueryProposalsParams
)

var (
	NewQueryProposalParams  = gov.NewQueryProposalParams
	NewQueryDepositParams   = gov.NewQueryDepositParams
	NewQueryVoteParams      = gov.NewQueryVoteParams
	NewQueryProposalsParams = gov.NewQueryProposalsParams
)
