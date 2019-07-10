package main

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/spf13/cobra"

	sdk "github.com/cosmos/cosmos-sdk/types"

	abci "github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tendermint/libs/db"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/tendermint/tendermint/libs/log"
	"github.com/tendermint/tmlibs/cli"

	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/x/auth/genaccounts"
	"github.com/cosmos/cosmos-sdk/x/staking"

	app "github.com/EG-easy/votumchain"

	genaccscli "github.com/cosmos/cosmos-sdk/x/auth/genaccounts/client/cli"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
)

const (
	flagOverwrite = "overwrite"
)

func main() {
	fmt.Println("start votumd!")
	cobra.EnableCommandSorting = false

	cdc := app.MakeCodec()

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(sdk.Bech32PrefixAccAddr, sdk.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(sdk.Bech32PrefixValAddr, sdk.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(sdk.Bech32PrefixConsAddr, sdk.Bech32PrefixConsPub)
	config.Seal()

	ctx := server.NewDefaultContext()

	rootCmd := &cobra.Command{
		Use:               "cotumchaind",
		Short:             "votumchain App Daemon (server)",
		PersistentPreRunE: server.PersistentPreRunEFn(ctx),
	}

	// CLI commands to initialize the chain
	rootCmd.AddCommand(
		genutilcli.InitCmd(ctx, cdc, app.ModuleBasics, app.DefaultNodeHome),
		genutilcli.CollectGenTxsCmd(ctx, cdc, genaccounts.AppModuleBasic{}, app.DefaultNodeHome),
		genutilcli.GenTxCmd(ctx, cdc, app.ModuleBasics, staking.AppModuleBasic{}, genaccounts.AppModuleBasic{}, app.DefaultNodeHome, app.DefaultCLIHome),
		genutilcli.ValidateGenesisCmd(ctx, cdc, app.ModuleBasics),
		// AddGenesisAccountCmd allows users to add accounts to the genesis file
		genaccscli.AddGenesisAccountCmd(ctx, cdc, app.DefaultNodeHome, app.DefaultCLIHome),
	)

	server.AddCommands(ctx, cdc, rootCmd, newApp, exportAppStateAndTMValidators)

	// prepare and add flags
	executor := cli.PrepareBaseCmd(rootCmd, "VOTUM", app.DefaultNodeHome)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

func newApp(logger log.Logger, db dbm.DB, traceStore io.Writer) abci.Application {
	return app.NewVotumApp(logger, db)
}

func exportAppStateAndTMValidators(
	logger log.Logger, db dbm.DB, traceStore io.Writer, height int64, forZeroHeight bool, jailWhiteList []string,
) (json.RawMessage, []tmtypes.GenesisValidator, error) {

	if height != -1 {
		votumApp := app.NewVotumApp(logger, db)
		err := votumApp.LoadHeight(height)
		if err != nil {
			return nil, nil, err
		}
		return votumApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
	}

	votumApp := app.NewVotumApp(logger, db)

	return votumApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
}
