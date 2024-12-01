package types

import (
	"cosmossdk.io/errors"
)

var (
	ErrRewardNotFound         = errors.Register(ModuleName, 1, "reward not found")
	ErrInsufficientRewards    = errors.Register(ModuleName, 2, "insufficient rewards to claim")
	ErrInvalidSlashAmount     = errors.Register(ModuleName, 3, "invalid slash amount")
	ErrNodeAlreadySlashed     = errors.Register(ModuleName, 4, "node already slashed for this event")
	ErrInvalidPerformanceData = errors.Register(ModuleName, 5, "invalid performance metrics")
	ErrRewardPoolEmpty        = errors.Register(ModuleName, 6, "reward pool is empty")
)

