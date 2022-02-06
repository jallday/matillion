package server

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/joshuaAllday/matillion/pkg/config"
)

func TestServer(t *testing.T) {
	s, err := New(config.LoadConfig())
	require.Nil(t, err)
	s.Start()
	s.Stop()
}
