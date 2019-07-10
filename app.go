package app

import (
	"encoding/json"
	"os"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/genaccounts"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	dbm "github.com/tendermint/tendermint/libs/db"
	tmtypes "github.com/tendermint/tendermint/types"

	distr "github.com/cosmos/cosmos-sdk/x/distribution"
)

const (
	appName = "votumApp"
)

var (
	// default home directories for votum app cli
	DefaultCLIHome = os.ExpandEnv("$HOME/.votumcli")

	// default home directories for gaiad
	DefaultNodeHome = os.ExpandEnv("$HOME/.votumd")

	// The ModuleBasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		genaccounts.AppModuleBasic{},
		genutil.AppModuleBasic{},
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		params.AppModuleBasic{},
		// TODO: add original module
		staking.AppModuleBasic{},
		distr.AppModuleBasic{},
		slashing.AppModuleBasic{},
	)
)

//cosmo-sdkのbaseAppを利用する
type votumApp struct {
	*bam.BaseApp
	cdc *codec.Codec

	// Keys to access the substores
	keyMain          *sdk.KVStoreKey
	keyAccount       *sdk.KVStoreKey
	keyFeeCollection *sdk.KVStoreKey
	keyStaking       *sdk.KVStoreKey
	tkeyStaking      *sdk.TransientStoreKey
	keyDistr         *sdk.KVStoreKey
	tkeyDistr        *sdk.TransientStoreKey
	keyParams        *sdk.KVStoreKey
	tkeyParams       *sdk.TransientStoreKey
	keySlashing      *sdk.KVStoreKey

	// Keepers
	accountKeeper       auth.AccountKeeper
	bankKeeper          bank.Keeper
	stakingKeeper       staking.Keeper
	slashingKeeper      slashing.Keeper
	distrKeeper         distr.Keeper
	feeCollectionKeeper auth.FeeCollectionKeeper
	paramsKeeper        params.Keeper
	// TODO add original keeper

	// Module Manager
	mm *module.Manager
}

//NewCoinTransferApp 新しいCoinTransactionAppの生成
func NewVotumApp(logger log.Logger, db dbm.DB) *votumApp {

	cdc := MakeCodec()

	//BaseAppは、ABCIを通してTendermintと相互にやりとりをする
	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc))

	var app = &votumApp{
		BaseApp: bApp,
		cdc:     cdc,

		keyMain:          sdk.NewKVStoreKey(bam.MainStoreKey),
		keyAccount:       sdk.NewKVStoreKey(auth.StoreKey),
		keyFeeCollection: sdk.NewKVStoreKey(auth.FeeStoreKey),
		keyStaking:       sdk.NewKVStoreKey(staking.StoreKey),
		tkeyStaking:      sdk.NewTransientStoreKey(staking.TStoreKey),
		keyDistr:         sdk.NewKVStoreKey(distr.StoreKey),
		tkeyDistr:        sdk.NewTransientStoreKey(distr.TStoreKey),
		keyParams:        sdk.NewKVStoreKey(params.StoreKey),
		tkeyParams:       sdk.NewTransientStoreKey(params.TStoreKey),
		keySlashing:      sdk.NewKVStoreKey(slashing.StoreKey),
	}

	// The ParamsKeeper handles parameter storage for the application
	app.paramsKeeper = params.NewKeeper(app.cdc, app.keyParams, app.tkeyParams, params.DefaultCodespace)
	// Set specific subspaces
	authSubspace := app.paramsKeeper.Subspace(auth.DefaultParamspace)
	bankSupspace := app.paramsKeeper.Subspace(bank.DefaultParamspace)
	stakingSubspace := app.paramsKeeper.Subspace(staking.DefaultParamspace)
	distrSubspace := app.paramsKeeper.Subspace(distr.DefaultParamspace)
	slashingSubspace := app.paramsKeeper.Subspace(slashing.DefaultParamspace)

	// The AccountKeeper handles address -> account lookups
	app.accountKeeper = auth.NewAccountKeeper(
		app.cdc,
		app.keyAccount,
		authSubspace,
		auth.ProtoBaseAccount,
	)

	// The BankKeeper allows you perform sdk.Coins interactions
	app.bankKeeper = bank.NewBaseKeeper(
		app.accountKeeper,
		bankSupspace,
		bank.DefaultCodespace,
	)

	// The FeeCollectionKeeper collects transaction fees and renders them to the fee distribution module
	app.feeCollectionKeeper = auth.NewFeeCollectionKeeper(cdc, app.keyFeeCollection)

	// The staking keeper
	stakingKeeper := staking.NewKeeper(
		app.cdc,
		app.keyStaking,
		app.tkeyStaking,
		app.bankKeeper,
		stakingSubspace,
		staking.DefaultCodespace,
	)

	app.distrKeeper = distr.NewKeeper(
		app.cdc,
		app.keyDistr,
		distrSubspace,
		app.bankKeeper,
		&stakingKeeper,
		app.feeCollectionKeeper,
		distr.DefaultCodespace,
	)

	app.slashingKeeper = slashing.NewKeeper(
		app.cdc,
		app.keySlashing,
		&stakingKeeper,
		slashingSubspace,
		slashing.DefaultCodespace,
	)
	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	app.stakingKeeper = *stakingKeeper.SetHooks(
		staking.NewMultiStakingHooks(
			app.distrKeeper.Hooks(),
			app.slashingKeeper.Hooks()),
	)

	// TODO:add original module

	app.mm = module.NewManager(
		genaccounts.NewAppModule(app.accountKeeper),
		genutil.NewAppModule(app.accountKeeper, app.stakingKeeper, app.BaseApp.DeliverTx),
		auth.NewAppModule(app.accountKeeper, app.feeCollectionKeeper),
		bank.NewAppModule(app.bankKeeper, app.accountKeeper),
		// TODO:add original module
		distr.NewAppModule(app.distrKeeper),
		slashing.NewAppModule(app.slashingKeeper, app.stakingKeeper),
		staking.NewAppModule(app.stakingKeeper, app.feeCollectionKeeper, app.distrKeeper, app.accountKeeper),
	)

	app.mm.SetOrderBeginBlockers(distr.ModuleName, slashing.ModuleName)
	app.mm.SetOrderEndBlockers(staking.ModuleName)

	// Sets the order of Genesis - Order matters, genutil is to always come last
	app.mm.SetOrderInitGenesis(
		genaccounts.ModuleName,
		distr.ModuleName,
		staking.ModuleName,
		auth.ModuleName,
		bank.ModuleName,
		slashing.ModuleName,
		// TODO:add original module
		genutil.ModuleName,
	)

	// register all module routes and module queriers
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter())

	// The initChainer handles translating the genesis.json file into initial state for the network
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)

	// The AnteHandler handles signature verification and transaction pre-processing
	app.SetAnteHandler(
		auth.NewAnteHandler(
			app.accountKeeper,
			app.feeCollectionKeeper,
			auth.DefaultSigVerificationGasConsumer,
		),
	)

	app.MountStores(
		app.keyMain,
		app.keyAccount,
		app.keyFeeCollection,
		app.keyStaking,
		app.tkeyStaking,
		app.keyDistr,
		app.tkeyDistr,
		app.keySlashing,
		// TODO : add original key
		app.keyParams,
		app.tkeyParams,
	)

	err := app.LoadLatestVersion(app.keyMain)
	if err != nil {
		cmn.Exit(err.Error())
	}

	return app

}

// GenesisState represents chain state at the start of the chain. Any initial state (account balances) are stored here.
type GenesisState map[string]json.RawMessage

func NewDefaultGenesisState() GenesisState {
	return ModuleBasics.DefaultGenesis()
}

func (app *votumApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState

	err := app.cdc.UnmarshalJSON(req.AppStateBytes, &genesisState)
	if err != nil {
		panic(err)
	}

	return app.mm.InitGenesis(ctx, genesisState)
}

func (app *votumApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

func (app *votumApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

func (app *votumApp) LoadHeight(height int64) error {
	return app.LoadVersion(height, app.keyMain)
}

//------------------------------------------------------------

func (app *votumApp) ExportAppStateAndValidators(forZeroHeight bool, jailWhiteList []string,
) (appState json.RawMessage, validators []tmtypes.GenesisValidator, err error) {

	// as if they could withdraw from the start of the next block
	ctx := app.NewContext(true, abci.Header{Height: app.LastBlockHeight()})

	genState := app.mm.ExportGenesis(ctx)
	appState, err = codec.MarshalJSONIndent(app.cdc, genState)
	if err != nil {
		return nil, nil, err
	}

	validators = staking.WriteValidators(ctx, app.stakingKeeper)

	return appState, validators, nil
}

// MakeCodec generates the necessary codecs for Amino
func MakeCodec() *codec.Codec {
	var cdc = codec.New()
	ModuleBasics.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	return cdc
}
