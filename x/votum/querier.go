package votum

import (
	"log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

//querierはapplicationの状態を確認するためのメソッドを定義する

const (
	QueryBalanceOf = "balanceOf"
)

//NewQuerier は、querierのroutingを行う
//queryには、Msgのようなinterfaceはないので、自分で場合分けする
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryBalanceOf:
			return queryBalanceOf(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown types query endpoint")
		}
	}
}

func queryBalanceOf(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	addr, error := sdk.AccAddressFromBech32(path[0])
	if error != nil {
		log.Fatal(error)
	}

	coins := keeper.coinKeeper.GetCoins(ctx, addr)

	if !coins.IsValid() {
		return []byte{}, sdk.ErrUnknownRequest("could not find coin")
	}

	return []byte(coins.String()), nil
}
