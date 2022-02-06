package sqlstore

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.com/joshuaAllday/matillion/internal/storetesting"
	"gitlab.com/joshuaAllday/matillion/pkg/utils/models"
)

func TestFilm(t *testing.T) {
	store, err := initStore(true)
	if err != nil {
		panic(err)
	}
	defer storetesting.CleanupTestingStore(store.settings)

	t.Run("GetFilm", func(t *testing.T) {
		testGetFilm(t, store)
	})

	t.Run("ListFilm", func(t *testing.T) {
		testListFilm(t, store)
	})
}

func testGetFilm(t *testing.T, store *SqlStore) {
	_, err := store.Film().Get(context.Background(), "2a8d52d1-bc27-410e-9c5a-a8d4c2278673")
	require.Nil(t, err)

	_, err = store.Film().Get(context.Background(), "2a8d52d1-bc27-410e-9c5a-a8d4c2278673asds")
	require.Error(t, err)
	var nEn *models.ErrNotFound
	if !errors.As(err, &nEn) {
		t.Error("error should be of type error not found")
	}
}

func testListFilm(t *testing.T, store *SqlStore) {
	films, err := store.Film().List(context.Background(), &models.FilmOptions{
		Page:    0,
		PerPage: 6,
	})
	require.Nil(t, err)
	assert.EqualValues(t, 6, len(films), "there should be 6 default films")

	films, err = store.Film().List(context.Background(), &models.FilmOptions{
		Page:    1,
		PerPage: 6,
	})
	require.Nil(t, err)
	assert.EqualValues(t, 0, len(films), "there are 6 default films so this should return 0 and an empty array")
}
