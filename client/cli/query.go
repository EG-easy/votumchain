package cli

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"fmt"
	"strings"
)

func GetCmdBalanceOf(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "balanceOf [addr]",
		Short: "balanceOf addr",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			addr := args[0]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/balanceOf/%s", queryRoute, addr), nil)
			if err != nil {
				fmt.Printf("could not find coin - %s \n", string(addr))
				return nil
			}

			var out balanceOfRes
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

type balanceOfRes struct {
	Value sdk.Coins `json:"value"`
}

func (w balanceOfRes) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Value: %s`, w.Value))
}
