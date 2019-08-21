package types

import "github.com/cosmos/cosmos-sdk/codec"

var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

//RegisterCodec codecでencode/decodeするためにtypesを登録する
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgIssueToken{}, "votum/issue", nil)
}
