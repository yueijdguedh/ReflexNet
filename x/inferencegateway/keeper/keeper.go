package keeper

import (
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/yueijdguedh/ReflexNet/x/inferencegateway/types"
)

type Keeper struct {
	cdc      codec.BinaryCodec
	storeKey storetypes.StoreKey
	memKey   storetypes.StoreKey
	authority string
}

func NewKeeper(cdc codec.BinaryCodec, storeKey storetypes.StoreKey, memKey storetypes.StoreKey, authority string) Keeper {
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

func (k Keeper) SetInferenceRequest(ctx sdk.Context, req types.InferenceRequest) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&req)
	store.Set(types.GetInferenceRequestKey(req.RequestId), bz)
}

func (k Keeper) GetInferenceRequest(ctx sdk.Context, requestID string) (types.InferenceRequest, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetInferenceRequestKey(requestID))
	if bz == nil {
		return types.InferenceRequest{}, false
	}
	var req types.InferenceRequest
	k.cdc.MustUnmarshal(bz, &req)
	return req, true
}

