package sqlstore

import (
	"context"
	"time"
)

type systemStore struct {
	*SqlStore
}

func newSystemStore(ss *SqlStore) *systemStore {
	return &systemStore{ss}
}

func (ss *systemStore) Ping(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	if err := ss.master().PingContext(ctx); err != nil {
		return err
	}

	if ss.replicaDBs != nil {
		for _, replica := range ss.replicaDBs {
			if err := replica.PingContext(ctx); err != nil {
				return err
			}
		}
	}

	return nil
}
