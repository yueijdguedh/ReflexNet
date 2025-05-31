package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type Params struct {
	MinStakeAmount sdk.Int
	MaxShardsPerNode uint32
}

func (p Params) Validate() error {
	return nil
}
