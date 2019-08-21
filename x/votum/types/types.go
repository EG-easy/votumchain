package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Tokens is a struct that contains all the token issued as vote right
type Issuers struct {
	Owner  sdk.AccAddress
	Tokens sdk.Coins
}

// Returns a new Issuers with the minprice as the price
func NewIssuers() Issuers {
	return Issuers{}
}

// implement fmt.Stringer
func (i Issuers) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
Tokens: %s`, i.Owner, i.Tokens))
}
