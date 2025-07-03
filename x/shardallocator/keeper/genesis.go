package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/yueijdguedh/ReflexNet/x/shardallocator/types"
)

func (k Keeper) InitGenesis(ctx sdk.Context, gs types.GenesisState) {
	k.SetParams(ctx, gs.Params)
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return &types.GenesisState{
		Params: k.GetParams(ctx),
	}
}
