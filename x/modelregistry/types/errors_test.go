package types_test

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestErrors(t *testing.T) {
	require.NotNil(t, ErrModelNotFound)
}
