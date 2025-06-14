package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

func (k Keeper) SplitFees(ctx sdk.Context, total sdk.Int) (model, nodes, treasury sdk.Int) {
	return total.QuoRaw(3), total.QuoRaw(3), total.QuoRaw(3)
}
