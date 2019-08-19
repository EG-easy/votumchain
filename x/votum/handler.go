package votum

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

/*
handlerは、Msgsを受け取った後に、行われるactionを定義する
MsgsのメソッドであるValidateBasicでは、Msgsのimput時点でのチェックを行うが、
ブロックチェーンを参照してチェックを行わないので、その確認作業はhandler側でやる必要がある
*/

//NewHandler はMsgのroutingを行う
func NewHandler(keeper bank.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgIssueToken:
			return handleMsgIssueToken(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized transfercoin Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

//IssueTokenのMsgを扱うためのHandler
func handleMsgIssueToken(ctx sdk.Context, keeper bank.Keeper, msg MsgIssueToken) sdk.Result {

	newCoin := sdk.NewCoin(msg.Coins[0].Denom, msg.Coins[0].Amount)
	issuer := msg.Owner

	coins := keeper.GetCoins(ctx, issuer)
	newCoins := sdk.NewCoins(newCoin).Add(coins)

	if ok := newCoins.IsValid(); !ok {
		return sdk.ErrInvalidCoins("Issuing New Coin is failed").Result()
	}
	if _, err := keeper.AddCoins(ctx, issuer, newCoins); err != nil {
		return sdk.ErrInvalidCoins("Issuing New Coin is failed").Result()
	}
	return sdk.Result{}
}
