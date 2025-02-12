package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) CalculateFee(ctx sdk.Context, computeUnits uint64) sdk.Int {
	params := k.GetParams(ctx)
	return params.BaseFee.Add(sdk.NewInt(int64(computeUnits)))
}
