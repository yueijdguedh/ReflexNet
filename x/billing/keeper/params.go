package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/yueijdguedh/ReflexNet/x/billing/types"
)

func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.DefaultParams()
}
