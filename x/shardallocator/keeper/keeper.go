package keeper

import (
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/yueijdguedh/ReflexNet/x/shardallocator/types"
)

type Keeper struct {
	cdc      codec.BinaryCodec
	storeKey storetypes.StoreKey
	memKey   storetypes.StoreKey
	authority string
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	memKey storetypes.StoreKey,
	authority string,
) Keeper {
	return Keeper{
		cdc:       cdc,
		storeKey:  storeKey,
		memKey:    memKey,
		authority: authority,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

func (k Keeper) GetAuthority() string {
	return k.authority
}

func (k Keeper) SetParams(ctx sdk.Context, params types.Params) error {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&params)
	store.Set(types.ParamsKey, bz)
	return nil
}

func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ParamsKey)
	if bz == nil {
		return types.DefaultParams()
	}
	var params types.Params
	k.cdc.MustUnmarshal(bz, &params)
	return params
}

func (k Keeper) SetNodeInfo(ctx sdk.Context, nodeInfo types.NodeInfo) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&nodeInfo)
	store.Set(types.GetNodeInfoKey(nodeInfo.Address), bz)
}

func (k Keeper) GetNodeInfo(ctx sdk.Context, address string) (types.NodeInfo, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetNodeInfoKey(address))
	if bz == nil {
		return types.NodeInfo{}, false
	}
	var nodeInfo types.NodeInfo
	k.cdc.MustUnmarshal(bz, &nodeInfo)
	return nodeInfo, true
}

func (k Keeper) SetShardAssignment(ctx sdk.Context, assignment types.ShardAssignment) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&assignment)
	store.Set(types.GetShardAssignmentKey(assignment.ModelId, assignment.ShardIndex), bz)
}

func (k Keeper) GetShardAssignment(ctx sdk.Context, modelID uint64, shardIndex uint32) (types.ShardAssignment, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetShardAssignmentKey(modelID, shardIndex))
	if bz == nil {
		return types.ShardAssignment{}, false
	}
	var assignment types.ShardAssignment
	k.cdc.MustUnmarshal(bz, &assignment)
	return assignment, true
}

