package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type Params struct {
	BaseFee sdk.Int
	ComputeUnitPrice sdk.Dec
}

func (p Params) Validate() error {
	return nil
}
