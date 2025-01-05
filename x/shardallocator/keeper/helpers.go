package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

func (k Keeper) IsNodeActive(ctx sdk.Context, address string) bool {
	nodeInfo, found := k.GetNodeInfo(ctx, address)
	return found && nodeInfo.Status == types.NodeStatus_NODE_STATUS_ACTIVE
}
