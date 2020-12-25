package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/EG-easy/votumchain/x/votum/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	votumTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	votumTxCmd.AddCommand(flags.PostCommands(
	// this line is used by starport scaffolding # 1
	)...)

	return votumTxCmd
}
