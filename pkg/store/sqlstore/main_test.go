package sqlstore

import (
	"gitlab.com/joshuaAllday/matillion/internal/storetesting"
)

func initStore(seed bool) (*SqlStore, error) {
	settings, err := storetesting.SetupTestingStore()
	if err != nil {
		return nil, err
	}
	return New(settings, seed)
}
