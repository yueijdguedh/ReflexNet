package types

import (
	"cosmossdk.io/errors"
)

var (
	ErrBillingRecordNotFound  = errors.Register(ModuleName, 1, "billing record not found")
	ErrInsufficientBalance    = errors.Register(ModuleName, 2, "insufficient balance")
	ErrInvalidFeeAmount       = errors.Register(ModuleName, 3, "invalid fee amount")
	ErrPaymentAlreadyCompleted = errors.Register(ModuleName, 4, "payment already completed")
	ErrInvalidDistribution    = errors.Register(ModuleName, 5, "invalid payment distribution")
)

