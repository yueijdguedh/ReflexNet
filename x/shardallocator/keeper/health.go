package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

func (k Keeper) CheckNodeHealth(ctx sdk.Context, nodeAddr string) bool {
	return true
}
