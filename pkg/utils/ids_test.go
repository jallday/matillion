package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIds(t *testing.T) {
	for i := 0; i < 100; i++ {
		id := NewID()
		require.True(t, ValidId(id))
		require.LessOrEqual(t, len(id), 36, "ids should be 36 char in length")
	}
}
