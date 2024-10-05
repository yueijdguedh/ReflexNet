package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/yueijdguedh/ReflexNet/x/modelregistry/types"
)

// InitGenesis initializes the module's state from a genesis state
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	// Set parameters
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}

	// Set next model ID
	k.SetNextModelID(ctx, genState.NextModelId)

	// Initialize models
	for _, model := range genState.Models {
		k.SetModel(ctx, model)
	}
}

// ExportGenesis exports the module's state to a genesis state
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return &types.GenesisState{
		Params:      k.GetParams(ctx),
		Models:      k.GetAllModels(ctx),
		NextModelId: k.GetNextModelID(ctx),
	}
}

