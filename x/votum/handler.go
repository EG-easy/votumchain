package votum

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
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
	newCoins := sdk.NewCoins(newCoin)

	acc := auth.NewBaseAccountWithAddress(msg.Owner)
	err := acc.SetCoins(newCoins)
	if err != nil {
		return sdk.ErrInvalidCoins("Issuing New Coin is failed").Result()
	}
	return sdk.Result{}
}

// func handleMsgBuyName(ctx sdk.Context, keeper Keeper, msg MsgBuyName) sdk.Result {
// 	if keeper.GetPrice(ctx, msg.Name).IsAllGT(msg.Bid) { // Checks if the the bid price is greater than the price paid by the current owner
// 		return sdk.ErrInsufficientCoins("Bid not high enough").Result() // If not, throw an error
// 	}
// 	if keeper.HasOwner(ctx, msg.Name) {
// 		err := keeper.coinKeeper.SendCoins(ctx, msg.Buyer, keeper.GetOwner(ctx, msg.Name), msg.Bid)
// 		if err != nil {
// 			return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
// 		}
// 	} else {
// 		_, err := keeper.coinKeeper.SubtractCoins(ctx, msg.Buyer, msg.Bid) // If so, deduct the Bid amount from the sender
// 		if err != nil {
// 			return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
// 		}
// 	}
// 	keeper.SetOwner(ctx, msg.Name, msg.Buyer)
// 	keeper.SetPrice(ctx, msg.Name, msg.Bid)
// 	return sdk.Result{}
// }
//
