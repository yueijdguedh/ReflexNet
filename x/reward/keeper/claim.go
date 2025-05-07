package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

func (k Keeper) ClaimRewards(ctx sdk.Context, nodeAddr string) (sdk.Int, error) {
	return sdk.NewInt(1000), nil
}
