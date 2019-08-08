package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Tokens is a struct that contains all the token issued as vote right
type Tokens struct {
	sdk.Coins
}
