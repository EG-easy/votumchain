package votum

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

//Keeper は、storageを管理するstructで、getter/setterのメソッドを実装する
type Keeper struct {
	coinKeeper bank.BaseKeeper //coinの発行にはKeeperより上位権限をもつBaseKeeperが必要
	cdc        *codec.Codec    //codecはgo-aminoを用いて、byte codeをdecode/encodeしてtendermint側と通信するときに必要となる
}

//NewKeeper 新しいKeeperを生成する
func NewKeeper(coinKeeper bank.BaseKeeper, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper: coinKeeper,
		cdc:        cdc,
	}
}
