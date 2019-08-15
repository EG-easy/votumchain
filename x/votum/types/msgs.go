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

//MsgIssueToken IssueTokenに関するMsg interfaceの実装
type MsgIssueToken struct {
	Owner sdk.AccAddress
	Coins sdk.Coins
}

//NewMsgIssueToken 新しいMsgIssueTokenを生成
func NewMsgIssueToken(owner sdk.AccAddress, coins sdk.Coins) MsgIssueToken {
	return MsgIssueToken{
		Owner: owner,
		Coins: coins,
	}
}

//Route module名を定義する
func (msg MsgIssueToken) Route() string {
	return "issuretoken"
}

//Type action名を決める
func (msg MsgIssueToken) Type() string {
	return "issure_token"
}

//ValidateBasic Msgsの中身のチェックをする
func (msg MsgIssueToken) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if msg.Coins[0].Amount.IsNegative() {
		return sdk.ErrUnknownRequest("Amount cannot be negative")
	}
	return nil
}

//GetSignBytes 署名するためのMsgデータを取得する
func (msg MsgIssueToken) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

//GetSigners 誰の署名が必要か定義する
func (msg MsgIssueToken) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}
