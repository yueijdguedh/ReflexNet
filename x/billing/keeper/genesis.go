package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/yueijdguedh/ReflexNet/x/billing/types"
)

func (k Keeper) InitGenesis(ctx sdk.Context, gs types.GenesisState) {
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return types.DefaultGenesis()
}
