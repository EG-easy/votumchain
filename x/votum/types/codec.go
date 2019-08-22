package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/gov"
)

var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

//RegisterCodec codecでencode/decodeするためにtypesを登録する
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgIssueToken{}, "votum/issue", nil)
	cdc.RegisterConcrete(gov.MsgSubmitProposal{}, "votum/MsgSubmitProposal", nil)
	cdc.RegisterConcrete(gov.MsgDeposit{}, "votum/MsgDeposit", nil)
	cdc.RegisterConcrete(gov.MsgVote{}, "votum/MsgVote", nil)
	cdc.RegisterConcrete(gov.TextProposal{}, "votum/TextProposal", nil)
}
