package app

import (
	"io"
	"os"
	"path/filepath"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/consensus"
	consensuskeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/yueijdguedh/ReflexNet/x/modelregistry"
	modelregistrykeeper "github.com/yueijdguedh/ReflexNet/x/modelregistry/keeper"
	modelregistrytypes "github.com/yueijdguedh/ReflexNet/x/modelregistry/types"
)

const (
	AppName = "reflexnet"
	Bech32MainPrefix = "reflex"
)

var (
	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		staking.AppModuleBasic{},
		consensus.AppModuleBasic{},
		modelregistry.AppModuleBasic{},
	)
)

// ReflexNetApp extends an ABCI application
type ReflexNetApp struct {
	*baseapp.BaseApp

	cdc               *codec.LegacyAmino
	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry
	txConfig          sdk.TxConfig

	// keys to access the substores
	keys    map[string]*storetypes.KVStoreKey
	tkeys   map[string]*storetypes.TransientStoreKey
	memKeys map[string]*storetypes.MemoryStoreKey

	// keepers
	AccountKeeper       authkeeper.AccountKeeper
	BankKeeper          bankkeeper.Keeper
	StakingKeeper       *stakingkeeper.Keeper
	ConsensusKeeper     consensuskeeper.Keeper
	ModelRegistryKeeper modelregistrykeeper.Keeper

	// module manager
	mm *module.Manager
}

// NewReflexNetApp returns a reference to an initialized ReflexNetApp.
func NewReflexNetApp(
	logger log.Logger,
	db corestore.KVStoreWithBatch,
	traceStore io.Writer,
	loadLatest bool,
	appOpts server.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) *ReflexNetApp {
	interfaceRegistry, err := types.NewInterfaceRegistryWithOptions(types.InterfaceRegistryOptions{
		ProtoFiles: proto.HybridResolver,
		SigningOptions: signing.Options{
			AddressCodec: address.NewBech32Codec(Bech32MainPrefix),
			ValidatorAddressCodec: address.NewBech32Codec(Bech32MainPrefix + "valoper"),
		},
	})
	if err != nil {
		panic(err)
	}
	appCodec := codec.NewProtoCodec(interfaceRegistry)
	legacyAmino := codec.NewLegacyAmino()
	txConfig := authtx.NewTxConfig(appCodec, authtx.DefaultSignModes)

	bApp := baseapp.NewBaseApp(AppName, logger, db, txConfig.TxDecoder(), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)
	bApp.SetTxEncoder(txConfig.TxEncoder())

	keys := storetypes.NewKVStoreKeys(
		authtypes.StoreKey,
		banktypes.StoreKey,
		stakingtypes.StoreKey,
		consensustypes.StoreKey,
		modelregistrytypes.StoreKey,
	)
	tkeys := storetypes.NewTransientStoreKeys()
	memKeys := storetypes.NewMemoryStoreKeys(modelregistrytypes.MemStoreKey)

	app := &ReflexNetApp{
		BaseApp:           bApp,
		cdc:               legacyAmino,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		txConfig:          txConfig,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
	}

	// Initialize keepers
	app.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[authtypes.StoreKey]),
		authtypes.ProtoBaseAccount,
		map[string][]string{
			banktypes.ModuleName: {authtypes.Minter, authtypes.Burner},
		},
		Bech32MainPrefix,
		authtypes.NewModuleAddress("gov").String(),
	)

	app.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[banktypes.StoreKey]),
		app.AccountKeeper,
		map[string]bool{},
		authtypes.NewModuleAddress("gov").String(),
		logger,
	)

	app.StakingKeeper = stakingkeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[stakingtypes.StoreKey]),
		app.AccountKeeper,
		app.BankKeeper,
		authtypes.NewModuleAddress("gov").String(),
		addresscodec.NewBech32Codec(Bech32MainPrefix + "valcons"),
		addresscodec.NewBech32Codec(Bech32MainPrefix + "valoper"),
	)

	app.ConsensusKeeper = consensuskeeper.NewKeeper(
		appCodec,
		runtime.NewKVStoreService(keys[consensustypes.StoreKey]),
		authtypes.NewModuleAddress("gov").String(),
		nil,
	)

	app.ModelRegistryKeeper = modelregistrykeeper.NewKeeper(
		appCodec,
		keys[modelregistrytypes.StoreKey],
		memKeys[modelregistrytypes.MemStoreKey],
		authtypes.NewModuleAddress("gov").String(),
	)

	// Create module manager
	app.mm = module.NewManager(
		auth.NewAppModule(appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts, nil),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper, nil),
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper, nil),
		consensus.NewAppModule(appCodec, app.ConsensusKeeper),
		modelregistry.NewAppModule(appCodec, app.ModelRegistryKeeper),
	)

	app.mm.SetOrderBeginBlockers(
		consensustypes.ModuleName,
		stakingtypes.ModuleName,
	)

	app.mm.SetOrderEndBlockers(
		stakingtypes.ModuleName,
	)

	app.mm.SetOrderInitGenesis(
		authtypes.ModuleName,
		banktypes.ModuleName,
		stakingtypes.ModuleName,
		consensustypes.ModuleName,
		modelregistrytypes.ModuleName,
	)

	app.mm.RegisterInvariants(nil)
	app.mm.RegisterServices(module.NewConfigurator(appCodec, bApp.MsgServiceRouter(), bApp.GRPCQueryRouter()))

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			panic(err)
		}
	}

	return app
}

// Name returns the name of the App
func (app *ReflexNetApp) Name() string { return app.BaseApp.Name() }

// LegacyAmino returns ReflexNetApp's amino codec.
func (app *ReflexNetApp) LegacyAmino() *codec.LegacyAmino {
	return app.cdc
}

// AppCodec returns ReflexNetApp's app codec.
func (app *ReflexNetApp) AppCodec() codec.Codec {
	return app.appCodec
}

// InterfaceRegistry returns ReflexNetApp's InterfaceRegistry
func (app *ReflexNetApp) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// TxConfig returns ReflexNetApp's TxConfig
func (app *ReflexNetApp) TxConfig() sdk.TxConfig {
	return app.txConfig
}

// GetKey returns the KVStoreKey for the provided store key.
func (app *ReflexNetApp) GetKey(storeKey string) *storetypes.KVStoreKey {
	return app.keys[storeKey]
}

// ModuleManager returns the app's module manager
func (app *ReflexNetApp) ModuleManager() *module.Manager {
	return app.mm
}

