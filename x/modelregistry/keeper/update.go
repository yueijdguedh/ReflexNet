package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/yueijdguedh/ReflexNet/x/modelregistry/types"
)

func (k Keeper) UpdateModel(ctx sdk.Context, model types.Model) error {
	k.SetModel(ctx, model)
	return nil
}
