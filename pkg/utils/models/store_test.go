package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStoreDefaults(t *testing.T) {
	s := &SqlSettings{}
	s.SetDefaults()
	require.Nil(t, s.ReplicaURLS)
	assert.EqualValues(t, MasterURL, s.MasterURL)
}
