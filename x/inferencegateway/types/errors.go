package types

import (
	"cosmossdk.io/errors"
)

var (
	ErrInvalidRequestID     = errors.Register(ModuleName, 1, "invalid request ID")
	ErrRequestNotFound      = errors.Register(ModuleName, 2, "inference request not found")
	ErrInvalidNonce         = errors.Register(ModuleName, 3, "invalid nonce")
	ErrRequestTimeout       = errors.Register(ModuleName, 4, "inference request timeout")
	ErrInvalidProof         = errors.Register(ModuleName, 5, "invalid zkML proof")
	ErrProofVerificationFailed = errors.Register(ModuleName, 6, "proof verification failed")
	ErrRateLimitExceeded    = errors.Register(ModuleName, 7, "rate limit exceeded")
	ErrRequestTooLarge      = errors.Register(ModuleName, 8, "request size exceeds maximum")
	ErrModelNotActive       = errors.Register(ModuleName, 9, "model is not active")
)

