package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

func (k Keeper) GetNonce(ctx sdk.Context, address string) uint64 {
	return 0
}

func (k Keeper) IncrementNonce(ctx sdk.Context, address string) {
}
