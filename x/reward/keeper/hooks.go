package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

type Hooks struct {
	k Keeper
}

func (k Keeper) Hooks() Hooks {
	return Hooks{k}
}
