package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:           DefaultParams(),
		ShardAssignments: []ShardAssignment{},
		NodeInfos:        []NodeInfo{},
	}
}

// DefaultParams returns default module parameters
func DefaultParams() Params {
	return Params{
		MinStakeAmount:       sdk.NewInt(10000000), // 10 MCELL
		MaxShardsPerNode:     50,
		HealthCheckInterval:  100, // blocks
		MinReputationScore:   0,
	}
}

// Validate performs basic validation of genesis data
func (gs GenesisState) Validate() error {
	// Validate params
	if err := gs.Params.Validate(); err != nil {
		return err
	}

	// Validate shard assignments
	assignmentKeys := make(map[string]bool)
	for _, assignment := range gs.ShardAssignments {
		if err := assignment.Validate(); err != nil {
			return fmt.Errorf("invalid shard assignment: %w", err)
		}

		key := fmt.Sprintf("%d-%d", assignment.ModelId, assignment.ShardIndex)
		if assignmentKeys[key] {
			return fmt.Errorf("duplicate shard assignment: model %d, shard %d", assignment.ModelId, assignment.ShardIndex)
		}
		assignmentKeys[key] = true
	}

	// Validate node infos
	nodeAddresses := make(map[string]bool)
	for _, nodeInfo := range gs.NodeInfos {
		if err := nodeInfo.Validate(); err != nil {
			return fmt.Errorf("invalid node info: %w", err)
		}

		if nodeAddresses[nodeInfo.Address] {
			return fmt.Errorf("duplicate node address: %s", nodeInfo.Address)
		}
		nodeAddresses[nodeInfo.Address] = true
	}

	return nil
}

// Validate performs validation on Params
func (p Params) Validate() error {
	if p.MinStakeAmount.IsNegative() {
		return fmt.Errorf("min stake amount cannot be negative")
	}

	if p.MaxShardsPerNode == 0 {
		return fmt.Errorf("max shards per node must be positive")
	}

	if p.HealthCheckInterval <= 0 {
		return fmt.Errorf("health check interval must be positive")
	}

	return nil
}

// Validate performs validation on ShardAssignment
func (sa ShardAssignment) Validate() error {
	if sa.ModelId == 0 {
		return fmt.Errorf("model ID cannot be zero")
	}

	if _, err := sdk.AccAddressFromBech32(sa.NodeAddress); err != nil {
		return fmt.Errorf("invalid node address: %w", err)
	}

	return nil
}

// Validate performs validation on NodeInfo
func (ni NodeInfo) Validate() error {
	if _, err := sdk.AccAddressFromBech32(ni.Address); err != nil {
		return fmt.Errorf("invalid node address: %w", err)
	}

	if ni.StakedAmount.IsNegative() {
		return fmt.Errorf("staked amount cannot be negative")
	}

	if ni.UptimePercentage.IsNegative() || ni.UptimePercentage.GT(sdk.OneDec()) {
		return fmt.Errorf("uptime percentage must be between 0 and 1")
	}

	return nil
}

