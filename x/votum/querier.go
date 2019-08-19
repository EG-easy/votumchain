package votum

import (
	"github.com/EG-easy/votumchain/x/votum/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	abci "github.com/tendermint/tendermint/abci/types"
)

//querierはapplicationの状態を確認するためのメソッドを定義する

const (
	QueryAccount = "account"
)

//NewQuerier は、querierのroutingを行う
//queryには、Msgのようなinterfaceはないので、自分で場合分けする
// creates a querier for auth REST endpoints
func NewQuerier(keeper bank.Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, sdk.Error) {
		switch path[0] {
		case types.QueryAccount:
			return queryAccount(ctx, req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown auth query endpoint")
		}
	}
}

func queryAccount(ctx sdk.Context, req abci.RequestQuery, keeper bank.Keeper) ([]byte, sdk.Error) {
	// var params types.QueryAccountParams
	// if err := keeper.cdc.UnmarshalJSON(req.Data, &params); err != nil {
	// 	return nil, sdk.ErrInternal(fmt.Sprintf("failed to parse params: %s", err))
	// }

	// account := keeper.ak.GetAccount(ctx, params.Address)
	// if account == nil {
	// 	return nil, sdk.ErrUnknownAddress(fmt.Sprintf("account %s does not exist", params.Address))
	// }

	// bz, err := codec.MarshalJSONIndent(keeper.cdc, account)
	// if err != nil {
	// 	return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", err.Error()))
	// }
	// return bz, nil

	return nil, nil
}
