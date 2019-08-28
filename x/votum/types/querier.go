package types

// // Params for queries:
// // - 'custom/gov/proposal'
// // - 'custom/gov/deposits'
// // - 'custom/gov/tally'
// // - 'custom/gov/votes'
// type QueryProposalParams struct {
// 	ProposalID uint64
// }
//
// // creates a new instance of QueryProposalParams
// func NewQueryProposalParams(proposalID uint64) QueryProposalParams {
// 	return QueryProposalParams{
// 		ProposalID: proposalID,
// 	}
// }
//
// // Params for query 'custom/gov/deposit'
// type QueryDepositParams struct {
// 	ProposalID uint64
// 	Depositor  sdk.AccAddress
// }
//
// // creates a new instance of QueryDepositParams
// func NewQueryDepositParams(proposalID uint64, depositor sdk.AccAddress) QueryDepositParams {
// 	return QueryDepositParams{
// 		ProposalID: proposalID,
// 		Depositor:  depositor,
// 	}
// }
//
// // Params for query 'custom/gov/vote'
// type QueryVoteParams struct {
// 	ProposalID uint64
// 	Voter      sdk.AccAddress
// }
//
// // creates a new instance of QueryVoteParams
// func NewQueryVoteParams(proposalID uint64, voter sdk.AccAddress) QueryVoteParams {
// 	return QueryVoteParams{
// 		ProposalID: proposalID,
// 		Voter:      voter,
// 	}
// }
//
// // Params for query 'custom/gov/proposals'
// type QueryProposalsParams struct {
// 	Voter          sdk.AccAddress
// 	Depositor      sdk.AccAddress
// 	ProposalStatus gov.ProposalStatus
// 	Limit          uint64
// }
//
// // creates a new instance of QueryProposalsParams
// func NewQueryProposalsParams(status gov.ProposalStatus, limit uint64, voter, depositor sdk.AccAddress) QueryProposalsParams {
// 	return QueryProposalsParams{
// 		Voter:          voter,
// 		Depositor:      depositor,
// 		ProposalStatus: status,
// 		Limit:          limit,
// 	}
// }
