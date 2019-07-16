package client

import (
	votumchaincmd "github.com/EG-easy/votumchain/client/cli"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	amino "github.com/tendermint/go-amino"
)

type ModuleClient struct {
	storeKey string
	cdc      *amino.Codec
}

func NewModuleClient(storeKey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient{storeKey, cdc}
}

func (mc ModuleClient) GetQueryCmd() *cobra.Command {
	votumchainQueryCmd := &cobra.Command{
		Use:   "votumchain",
		Short: "Querying commands for the votumchain module",
	}

	votumchainQueryCmd.AddCommand(client.GetCommands(
		votumchaincmd.GetCmdBalanceOf(mc.storeKey, mc.cdc),
	)...)

	return votumchainQueryCmd
}

func (mc ModuleClient) GetTxCmd() *cobra.Command {
	votumchainTxCmd := &cobra.Command{
		Use:   "votumchain",
		Short: "votumchain transactions subcommands",
	}

	votumchainTxCmd.AddCommand(client.PostCommands(
		votumchaincmd.GetCmdTransfer(mc.cdc),
	)...)

	return votumchainTxCmd
}
