package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/yueijdguedh/ReflexNet/x/billing/types"
)

func (k Keeper) CreateBillingRecord(ctx sdk.Context, record types.BillingRecord) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&record)
	store.Set(types.GetBillingRecordKey(record.RequestId), bz)
}
