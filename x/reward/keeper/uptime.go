package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

func (k Keeper) CalculateUptime(ctx sdk.Context, nodeAddr string) sdk.Dec {
	return sdk.NewDec(98).QuoInt64(100)
}
