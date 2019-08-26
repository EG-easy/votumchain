package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

//RegisterCodec codecでencode/decodeするためにtypesを登録する
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterInterface((*Content)(nil), nil)

	cdc.RegisterConcrete(MsgIssueToken{}, "votum/issue", nil)
	cdc.RegisterConcrete(MsgSubmitProposal{}, "votum/MsgSubmitProposal", nil)
	cdc.RegisterConcrete(MsgDeposit{}, "votum/MsgDeposit", nil)
	cdc.RegisterConcrete(MsgVote{}, "votum/MsgVote", nil)
	cdc.RegisterConcrete(TextProposal{}, "votum/TextProposal", nil)
}

// // RegisterProposalTypeCodec registers an external proposal content type defined
// // in another module for the internal ModuleCdc. This allows the MsgSubmitProposal
// // to be correctly Amino encoded and decoded.
// func RegisterProposalTypeCodec(o interface{}, name string) {
// 	ModuleCdc.RegisterConcrete(o, name, nil)
// }
