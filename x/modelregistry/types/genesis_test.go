package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/yueijdguedh/ReflexNet/x/modelregistry/types"
)

func TestDefaultGenesis(t *testing.T) {
	genesis := types.DefaultGenesis()
	require.NotNil(t, genesis)
	require.NotNil(t, genesis.Params)
	require.Equal(t, uint64(1), genesis.NextModelId)
	require.Empty(t, genesis.Models)
}

func TestGenesisValidation(t *testing.T) {
	testCases := []struct {
		name      string
		genesis   *types.GenesisState
		expectErr bool
	}{
		{
			name:      "valid default genesis",
			genesis:   types.DefaultGenesis(),
			expectErr: false,
		},
		{
			name: "invalid next model ID",
			genesis: &types.GenesisState{
				Params:      types.DefaultParams(),
				Models:      []types.Model{},
				NextModelId: 0,
			},
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.genesis.Validate()
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestParamsValidation(t *testing.T) {
	testCases := []struct {
		name      string
		params    types.Params
		expectErr bool
	}{
		{
			name:      "valid default params",
			params:    types.DefaultParams(),
			expectErr: false,
		},
		{
			name: "negative registration fee",
			params: types.Params{
				RegistrationFee: sdk.NewInt(-1),
				MaxMetadataSize: 1024,
				MaxShardCount:   10,
			},
			expectErr: true,
		},
		{
			name: "zero max metadata size",
			params: types.Params{
				RegistrationFee: sdk.NewInt(1000),
				MaxMetadataSize: 0,
				MaxShardCount:   10,
			},
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.params.Validate()
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

