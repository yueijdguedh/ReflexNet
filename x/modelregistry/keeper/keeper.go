package keeper

import (
	"fmt"

	"cosmossdk.io/log"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/yueijdguedh/ReflexNet/x/modelregistry/types"
)

type Keeper struct {
	cdc      codec.BinaryCodec
	storeKey storetypes.StoreKey
	memKey   storetypes.StoreKey

	authority string
}

// NewKeeper creates a new modelregistry Keeper instance
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

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// SetParams sets the module parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) error {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&params)
	store.Set(types.ParamsKey, bz)
	return nil
}

// GetParams returns the module parameters.
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

// GetNextModelID returns the next model ID and increments it
func (k Keeper) GetNextModelID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.NextModelIDKey)
	if bz == nil {
		return 1
	}
	return types.BytesToUint64(bz)
}

// SetNextModelID sets the next model ID
func (k Keeper) SetNextModelID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.NextModelIDKey, types.Uint64ToBytes(id))
}

// SetModel stores a model in the store
func (k Keeper) SetModel(ctx sdk.Context, model types.Model) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&model)
	store.Set(types.GetModelKey(model.Id), bz)

	// Index by owner
	ownerStore := prefix.NewStore(store, types.ModelByOwnerKey)
	ownerStore.Set(append([]byte(model.Owner), types.Uint64ToBytes(model.Id)...), []byte{})
}

// GetModel retrieves a model by ID
func (k Keeper) GetModel(ctx sdk.Context, id uint64) (types.Model, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetModelKey(id))
	if bz == nil {
		return types.Model{}, false
	}

	var model types.Model
	k.cdc.MustUnmarshal(bz, &model)
	return model, true
}

// DeleteModel removes a model from the store
func (k Keeper) DeleteModel(ctx sdk.Context, id uint64) {
	model, found := k.GetModel(ctx, id)
	if !found {
		return
	}

	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetModelKey(id))

	// Remove from owner index
	ownerStore := prefix.NewStore(store, types.ModelByOwnerKey)
	ownerStore.Delete(append([]byte(model.Owner), types.Uint64ToBytes(model.Id)...))
}

// GetAllModels returns all models
func (k Keeper) GetAllModels(ctx sdk.Context) []types.Model {
	store := ctx.KVStore(k.storeKey)
	iterator := storetypes.KVStorePrefixIterator(store, types.ModelKey)
	defer iterator.Close()

	var models []types.Model
	for ; iterator.Valid(); iterator.Next() {
		var model types.Model
		k.cdc.MustUnmarshal(iterator.Value(), &model)
		models = append(models, model)
	}
	return models
}

// GetModelsByOwner returns all models owned by a specific address
func (k Keeper) GetModelsByOwner(ctx sdk.Context, owner string) []types.Model {
	store := ctx.KVStore(k.storeKey)
	ownerStore := prefix.NewStore(store, types.ModelByOwnerKey)
	iterator := storetypes.KVStorePrefixIterator(ownerStore, []byte(owner))
	defer iterator.Close()

	var models []types.Model
	for ; iterator.Valid(); iterator.Next() {
		// Extract model ID from the key
		key := iterator.Key()
		modelID := types.BytesToUint64(key[len(owner):])
		
		if model, found := k.GetModel(ctx, modelID); found {
			models = append(models, model)
		}
	}
	return models
}

