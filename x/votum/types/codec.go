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
	cdc.RegisterConcrete(MsgIssueToken{}, "votum/issue", nil)
	cdc.RegisterConcrete(MsgSubmitProposal{}, "votum/MsgSubmitProposal", nil)
	cdc.RegisterConcrete(MsgDeposit{}, "votum/MsgDeposit", nil)
	cdc.RegisterConcrete(MsgVote{}, "votum/MsgVote", nil)
	cdc.RegisterConcrete(TextProposal{}, "votum/TextProposal", nil)
}
