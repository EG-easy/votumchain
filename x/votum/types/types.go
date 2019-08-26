package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
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

// Proposal types
const (
	ProposalTypeText string = "Text"
)

// Text Proposal
type TextProposal struct {
	Title       string `json:"title" yaml:"title"`
	Description string `json:"description" yaml:"description"`
}

func NewTextProposal(title, description string) Content {
	return TextProposal{title, description}
}

// Implements Proposal Interface
var _ Content = TextProposal{}

// nolint
func (tp TextProposal) GetTitle() string         { return tp.Title }
func (tp TextProposal) GetDescription() string   { return tp.Description }
func (tp TextProposal) ProposalRoute() string    { return RouterKey }
func (tp TextProposal) ProposalType() string     { return ProposalTypeText }
func (tp TextProposal) ValidateBasic() sdk.Error { return gov.ValidateAbstract(DefaultCodespace, tp) }

func (tp TextProposal) String() string {
	return fmt.Sprintf(`Text Proposal:
  Title:       %s
  Description: %s
`, tp.Title, tp.Description)
}

// ContentFromProposalType returns a Content object based on the proposal type.
func ContentFromProposalType(title, desc, ty string) Content {
	switch ty {
	case ProposalTypeText:
		return NewTextProposal(title, desc)
	default:
		return nil
	}
}
