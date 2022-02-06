package sqlstore

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/joshuaAllday/matillion/internal/storetesting"
)

func TestSystem(t *testing.T) {
	store, err := initStore(false)
	if err != nil {
		panic(err)
	}
	defer storetesting.CleanupTestingStore(store.settings)

	t.Run("SystemPing", func(t *testing.T) {
		testSystemPing(t, store)
	})
}

func testSystemPing(t *testing.T, store *SqlStore) {
	defer func() { require.Nil(t, store.truncateTables()) }()
	err := store.System().Ping(context.Background())
	require.Nil(t, err)
}
