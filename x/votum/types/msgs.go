package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName // this was defined in your key.go file

/*
Msgsは、stateの変更のtransactionを実行する時に渡すパラメータで、Txsにラップされてnetworkにbroadcastされる
cosmos-sdkを使うことで、Msgsをラップしたり、Txsからのアンラップしたりしてくれるので、実質的にMsgsの定義だけすれば良い
Msgsは、次のinterfaceを満たすように実装すれば良い。

type Msg interface {
	Type() string
	Route() string
	ValidateBasic() Error
	GetSignBytes() []byte
	GetSigners() []AccAddress
}
*/

//MsgTransferCoin TransferCoinに関するMsg interfaceの実装
type MsgTransferCoin struct {
	FromAddr sdk.AccAddress
	ToAddr   sdk.AccAddress
	Amt      sdk.Coins
}

//NewMsgTransferCoin 新しいMsgTransferCoinを生成
func NewMsgTransferCoin(fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) MsgTransferCoin {
	return MsgTransferCoin{
		FromAddr: fromAddr,
		ToAddr:   toAddr,
		Amt:      amt,
	}
}

//Route module名を定義する
func (msg MsgTransferCoin) Route() string {
	return "transfercoin"
}

//Type action名を決める
func (msg MsgTransferCoin) Type() string {
	return "transfer_coin"
}

//ValidateBasic Msgsの中身のチェックをする
func (msg MsgTransferCoin) ValidateBasic() sdk.Error {
	if msg.ToAddr.Empty() {
		return sdk.ErrInvalidAddress(msg.ToAddr.String())
	}
	if len(msg.Amt) == 0 {
		return sdk.ErrUnknownRequest("Amount cannot be empty")
	}
	return nil
}

//GetSignBytes 署名するためのMsgデータを取得する
func (msg MsgTransferCoin) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

//GetSigners 誰の署名が必要か定義する
func (msg MsgTransferCoin) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.FromAddr}
}
