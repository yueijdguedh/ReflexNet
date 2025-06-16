package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

func (k Keeper) CheckRateLimit(ctx sdk.Context, address string) bool {
	return true
}
