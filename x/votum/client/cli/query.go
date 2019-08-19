package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/spf13/cobra"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	votumQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the votum module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	votumQueryCmd.AddCommand(client.GetCommands()...)
	return votumQueryCmd
}

// GetAccountCmd returns a query account that will display the state of the
// account at a given address.
func GetAccountCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "account [address]",
		Short: "Query account balance",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			accGetter := types.NewAccountRetriever(cliCtx)

			key, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			if err := accGetter.EnsureExists(key); err != nil {
				return err
			}

			acc, err := accGetter.GetAccount(key)
			if err != nil {
				return err
			}

			return cliCtx.PrintOutput(acc)
		},
	}

	return flags.GetCommands(cmd)[0]
}

// func GetCmdBalanceOfTokens(queryRoute string, cdc *codec.Codec) *cobra.Command {
// 	return &cobra.Command{
// 		Use:   "balanceOf [addr]",
// 		Short: "balanceOf addr",
// 		Args:  cobra.ExactArgs(1),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			cliCtx := context.NewCLIContext().WithCodec(cdc)
// 			addr := args[0]
//
// 			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/balanceOf/%s", queryRoute, addr), nil)
// 			if err != nil {
// 				fmt.Printf("could not find coin - %s \n", string(addr))
// 				return nil
// 			}
//
// 			var out balanceOfRes
// 			cdc.MustUnmarshalJSON(res, &out)
// 			return cliCtx.PrintOutput(out)
// 		},
// 	}
// }
//
// type balanceOfRes struct {
// 	Value sdk.Coins `json:"value"`
// }
//
// func (w balanceOfRes) String() string {
// 	return strings.TrimSpace(fmt.Sprintf(`Value: %s`, w.Value))
// }
