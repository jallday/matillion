package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSystemToJSON(t *testing.T) {
	require.NotPanics(t, func() {
		s := &System{Status: "Success"}
		assert.NotEqual(t, "", s.ToJSON(), "function should happily json marshal the struct into a non empty string")
	})

}
