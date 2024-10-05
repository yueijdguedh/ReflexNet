package types

import (
	"cosmossdk.io/errors"
)

// x/modelregistry module sentinel errors
var (
	ErrModelNotFound         = errors.Register(ModuleName, 1, "model not found")
	ErrInvalidModelID        = errors.Register(ModuleName, 2, "invalid model ID")
	ErrUnauthorized          = errors.Register(ModuleName, 3, "unauthorized")
	ErrInvalidMetadataCID    = errors.Register(ModuleName, 4, "invalid metadata CID")
	ErrInvalidShardCount     = errors.Register(ModuleName, 5, "invalid shard count")
	ErrModelNameTooLong      = errors.Register(ModuleName, 6, "model name too long")
	ErrInvalidVersion        = errors.Register(ModuleName, 7, "invalid version string")
	ErrInsufficientFee       = errors.Register(ModuleName, 8, "insufficient registration fee")
	ErrMetadataTooLarge      = errors.Register(ModuleName, 9, "metadata size exceeds maximum")
	ErrShardCountExceedsMax  = errors.Register(ModuleName, 10, "shard count exceeds maximum")
	ErrModelAlreadyExists    = errors.Register(ModuleName, 11, "model already exists")
	ErrInvalidModelStatus    = errors.Register(ModuleName, 12, "invalid model status")
)

