package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

/*
handlerは、Msgsを受け取った後に、行われるactionを定義する
MsgsのメソッドであるValidateBasicでは、Msgsのimput時点でのチェックを行うが、
ブロックチェーンを参照してチェックを行わないので、その確認作業はhandler側でやる必要がある
*/

//NewHandler はMsgのroutingを行う
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgTransferCoin:
			return handleMsgTransferCoin(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized transfercoin Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

//TransferCoinのMsgを扱うためのHandler
func handleMsgTransferCoin(ctx sdk.Context, keeper Keeper, msg MsgTransferCoin) sdk.Result {

	_, err := keeper.coinKeeper.SendCoins(ctx, msg.FromAddr, msg.ToAddr, msg.Amt)
	if err != nil {
		return sdk.ErrInsufficientCoins("Sender does not have enough coins").Result()
	}

	return sdk.Result{}
}
