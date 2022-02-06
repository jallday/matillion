package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.com/joshuaAllday/matillion/pkg/utils"
)

func TestFilm(t *testing.T) {
	require.NotPanics(t, func() {
		filmRating := &FilmRatings{Film: &Film{ID: utils.NewID()}}
		assert.NotEqual(t, "", filmRating.ToJSON(), "function should happily json marshal the struct into a non empty string")
	})

}
