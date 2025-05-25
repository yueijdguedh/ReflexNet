package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/yueijdguedh/ReflexNet/x/billing/types"
)

func (k Keeper) GetBillingRecord(ctx sdk.Context, requestID string) (types.BillingRecord, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetBillingRecordKey(requestID))
	if bz == nil {
		return types.BillingRecord{}, false
	}
	var record types.BillingRecord
	k.cdc.MustUnmarshal(bz, &record)
	return record, true
}
