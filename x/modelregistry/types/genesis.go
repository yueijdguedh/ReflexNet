package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:      DefaultParams(),
		Models:      []Model{},
		NextModelId: 1,
	}
}

// DefaultParams returns default module parameters
func DefaultParams() Params {
	return Params{
		RegistrationFee:  sdk.NewInt(1000000), // 1 MCELL
		MaxMetadataSize:  1024 * 1024,         // 1MB
		MaxShardCount:    100,
	}
}

// Validate performs basic validation of genesis data
func (gs GenesisState) Validate() error {
	// Validate params
	if err := gs.Params.Validate(); err != nil {
		return err
	}

	// Validate models
	modelIDs := make(map[uint64]bool)
	for _, model := range gs.Models {
		if err := model.Validate(); err != nil {
			return fmt.Errorf("invalid model with ID %d: %w", model.Id, err)
		}

		if modelIDs[model.Id] {
			return fmt.Errorf("duplicate model ID: %d", model.Id)
		}
		modelIDs[model.Id] = true
	}

	// Validate next model ID
	if gs.NextModelId < 1 {
		return fmt.Errorf("next model ID must be at least 1, got: %d", gs.NextModelId)
	}

	return nil
}

// Validate performs validation on Params
func (p Params) Validate() error {
	if p.RegistrationFee.IsNegative() {
		return fmt.Errorf("registration fee cannot be negative")
	}

	if p.MaxMetadataSize == 0 {
		return fmt.Errorf("max metadata size must be positive")
	}

	if p.MaxShardCount == 0 {
		return fmt.Errorf("max shard count must be positive")
	}

	return nil
}

// Validate performs validation on Model
func (m Model) Validate() error {
	if m.Id == 0 {
		return fmt.Errorf("model ID cannot be zero")
	}

	if _, err := sdk.AccAddressFromBech32(m.Owner); err != nil {
		return fmt.Errorf("invalid owner address: %w", err)
	}

	if m.Name == "" {
		return fmt.Errorf("model name cannot be empty")
	}

	if len(m.Name) > 256 {
		return ErrModelNameTooLong
	}

	if m.MetadataCid == "" {
		return ErrInvalidMetadataCID
	}

	if m.ShardCount == 0 {
		return ErrInvalidShardCount
	}

	if m.Version == "" {
		return ErrInvalidVersion
	}

	return nil
}

