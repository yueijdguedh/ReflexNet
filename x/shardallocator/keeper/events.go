package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

func EmitNodeRegisteredEvent(ctx sdk.Context, nodeAddr string) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent("node_registered"),
	)
}
