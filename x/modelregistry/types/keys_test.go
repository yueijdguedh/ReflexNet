package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yueijdguedh/ReflexNet/x/modelregistry/types"
)

func TestGetModelKey(t *testing.T) {
	key := types.GetModelKey(1)
	require.NotNil(t, key)
	require.Greater(t, len(key), 0)
}

func TestUint64Conversion(t *testing.T) {
	testCases := []uint64{0, 1, 100, 999999, 18446744073709551615}
	
	for _, tc := range testCases {
		bytes := types.Uint64ToBytes(tc)
		result := types.BytesToUint64(bytes)
		require.Equal(t, tc, result, "conversion failed for %d", tc)
	}
}

func TestGetModelByOwnerKey(t *testing.T) {
	owner := "reflex1abc123def456"
	modelID := uint64(42)
	
	key := types.GetModelByOwnerKey(owner, modelID)
	require.NotNil(t, key)
	require.Greater(t, len(key), len(owner))
}

