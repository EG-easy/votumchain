package types

import "github.com/cosmos/cosmos-sdk/codec"

//RegisterCodec codecでencode/decodeするためにtypesを登録する
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgTransferCoin{}, "types/TransferCoin", nil)
}
