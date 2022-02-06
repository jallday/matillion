package sqlstore

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.com/joshuaAllday/matillion/internal/storetesting"
	"gitlab.com/joshuaAllday/matillion/pkg/utils/models"
)

func TestRatings(t *testing.T) {
	store, err := initStore(true)
	if err != nil {
		panic(err)
	}
	defer storetesting.CleanupTestingStore(store.settings)

	t.Run("CreateRatings", func(t *testing.T) {
		testCreateRatings(t, store)
	})
	t.Run("ListRatings", func(t *testing.T) {
		testListRatings(t, store)
	})
}

func testCreateRatings(t *testing.T, store *SqlStore) {
	defer func() { require.Nil(t, store.truncateTables()) }()
	rating := &models.Rating{
		Author: "TestAuthor",
		Score:  1,
		FilmID: "2a8d52d1-bc27-410e-9c5a-a8d4c2278673",
	}

	rating, err := store.Ratings().Save(context.Background(), rating)
	require.Nil(t, err)

	_, err = store.Ratings().Save(context.Background(), rating)
	require.Error(t, err, "this should error and due an author rating this film already")
	var nErc *models.ErrConflict
	if !errors.As(err, &nErc) {
		t.Error("this should be a conflict error")
	}

	rating.Score = 10
	_, err = store.Ratings().Save(context.Background(), rating)
	require.Error(t, err, "this should error and due to invalid score")
	var nErr *models.Error
	if !errors.As(err, &nErr) {
		t.Error("this should be a type of *models.Error")
	}
}

func testListRatings(t *testing.T, store *SqlStore) {
	defer func() { require.Nil(t, store.truncateTables()) }()

	for i := 0; i < 10; i++ {
		rating := &models.Rating{
			Author: fmt.Sprintf("TestAuthor-%v", i),
			Score:  1,
			FilmID: "2a8d52d1-bc27-410e-9c5a-a8d4c2278673",
		}
		_, err := store.Ratings().Save(context.Background(), rating)
		require.Nil(t, err)
	}

	ratings, err := store.Ratings().List(context.Background(), &models.RatingOptions{
		Id:       "2a8d52d1-bc27-410e-9c5a-a8d4c2278673",
		Page:     0,
		PerPage:  10,
		MinScore: 0,
		MaxScore: 5,
	})
	require.Nil(t, err)
	assert.EqualValues(t, 10, len(ratings), "there should be 10 ratings between 0-5 for this film")

	ratings, err = store.Ratings().List(context.Background(), &models.RatingOptions{
		Id:       "2a8d52d1-bc27-410e-9c5a-a8d4c2278673",
		Page:     0,
		PerPage:  10,
		MinScore: 3,
		MaxScore: 5,
	})
	require.Nil(t, err)
	assert.EqualValues(t, 0, len(ratings), "there should be 0 ratings between 3-5 for this film")
}
