package types

import (
	"cosmossdk.io/errors"
)

// x/shardallocator module sentinel errors
var (
	ErrNodeNotFound           = errors.Register(ModuleName, 1, "node not found")
	ErrShardNotFound          = errors.Register(ModuleName, 2, "shard assignment not found")
	ErrInsufficientStake      = errors.Register(ModuleName, 3, "insufficient stake amount")
	ErrMaxShardsExceeded      = errors.Register(ModuleName, 4, "maximum shards per node exceeded")
	ErrNodeAlreadyRegistered  = errors.Register(ModuleName, 5, "node already registered")
	ErrNodeNotActive          = errors.Register(ModuleName, 6, "node is not active")
	ErrInvalidReputationScore = errors.Register(ModuleName, 7, "invalid reputation score")
	ErrUnauthorized           = errors.Register(ModuleName, 8, "unauthorized")
	ErrInvalidNodeStatus      = errors.Register(ModuleName, 9, "invalid node status")
	ErrShardAlreadyAssigned   = errors.Register(ModuleName, 10, "shard already assigned")
)

