package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

//Keeper は、storageを管理するstructで、getter/setterのメソッドを実装する
type Keeper struct {
	coinKeeper bank.Keeper //coinの送受信等に必要なモジュール

	storeKey sdk.StoreKey

	cdc *codec.Codec //codecはgo-aminoを用いて、byte codeをdecode/encodeしてtendermint側と通信するときに必要となる
}

//NewKeeper 新しいKeeperを生成する
func NewKeeper(coinKeeper bank.Keeper, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper: coinKeeper,
		cdc:        cdc,
	}
}
