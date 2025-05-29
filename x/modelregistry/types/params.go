package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type Params struct {
	RegistrationFee sdk.Int
	MaxMetadataSize uint64
	MaxShardCount uint32
}

func (p Params) Validate() error {
	return nil
}
