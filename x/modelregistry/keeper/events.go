package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

func EmitModelRegisteredEvent(ctx sdk.Context, modelID uint64) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent("model_registered"),
	)
}
