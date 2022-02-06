package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadingConfig(t *testing.T) {
	require.NotPanics(t, func() {
		LoadConfig()
	})
}
