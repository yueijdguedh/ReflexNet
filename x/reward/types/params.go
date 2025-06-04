package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type Params struct {
	BaseRewardAmount sdk.Int
}

func (p Params) Validate() error {
	return nil
}
