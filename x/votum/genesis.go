package votum

import (
	"fmt"

	"github.com/EG-easy/votumchain/x/votum/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	IssuerRecords []types.Issuers `json:"tokens"`
}

func NewGenesisState(issuerRecords []types.Issuers) GenesisState {
	return GenesisState{IssuerRecords: nil}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.IssuerRecords {
		if record.Tokens == nil {
			return fmt.Errorf("Invalid WhoisRecord: Value: %s. Error: Missing Owner", record.Tokens)
		}
	}
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		IssuerRecords: []types.Issuers{},
	}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	// for _, record := range data.IssuerRecords {
	// 	keeper.SetWhois(ctx, record.Value, record)
	// }
	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	// var records []Whois
	// iterator := k.GetNamesIterator(ctx)
	// for ; iterator.Valid(); iterator.Next() {
	// 	name := string(iterator.Key())
	// 	var whois Whois
	// 	whois = k.GetWhois(ctx, name)
	// 	records = append(records, whois)
	// }
	return GenesisState{
		IssuerRecords: []types.Issuers{},
	}
}
