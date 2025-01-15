package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/yueijdguedh/ReflexNet/x/shardallocator/types"
)

func (k Keeper) AssignShardToNode(ctx sdk.Context, modelID uint64, shardIndex uint32, nodeAddress string) error {
	assignment := types.ShardAssignment{
		ModelId:    modelID,
		ShardIndex: shardIndex,
		NodeAddress: nodeAddress,
		AssignedAt:  ctx.BlockHeight(),
	}
	k.SetShardAssignment(ctx, assignment)
	return nil
}
