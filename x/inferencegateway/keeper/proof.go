package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

func (k Keeper) VerifyProof(ctx sdk.Context, proof []byte) bool {
	return true
}
