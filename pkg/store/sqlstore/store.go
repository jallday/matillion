package sqlstore

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"gitlab.com/joshuaAllday/matillion/pkg/store"
	"gitlab.com/joshuaAllday/matillion/pkg/utils/models"
)

type SqlStore struct {
	replicaId  int
	masterDB   *sql.DB
	replicaDBs []*sql.DB
	stores     *stores
	settings   *models.SqlSettings
}

type stores struct {
	system  *systemStore
	films   *filmStore
	ratings *ratingsStore
}

func New(settings *models.SqlSettings, seed bool) (*SqlStore, error) {

	ss := &SqlStore{
		settings: settings,
	}

	if err := ss.loadConnections(); err != nil {
		return nil, err
	}

	ss.loadStores()
	if err := ss.loadSchema(); err != nil {
		return nil, err
	}

	if seed {
		if err := ss.loadSeeds(); err != nil {
			fmt.Printf("problem with loading seeds: %s\n", err.Error())
		}
	}

	// TODO(Josh): here we can look to add migration functions :)

	return ss, nil
}

func (ss *SqlStore) loadConnections() error {
	masterDB, err := connect(ss.settings.MasterURL, ss.settings, MasterFlag)
	if err != nil {
		return err
	}
	ss.masterDB = masterDB

	if len(ss.settings.ReplicaURLS) > 0 {
		replicas := make([]*sql.DB, 0)
		for _, url := range ss.settings.ReplicaURLS {
			replica, err := connect(url, ss.settings, ReplicaFlag)
			if err != nil {
				return err
			}
			replicas = append(replicas, replica)
		}
		ss.replicaDBs = replicas

	}

	return nil
}

func (ss *SqlStore) getTables() ([]string, error) {
	var tables []string
	res, _ := ss.replica().Query(`
	SELECT tablename
	FROM pg_catalog.pg_tables
	WHERE schemaname != 'pg_catalog' AND 
    schemaname != 'information_schema';
	`)

	for res.Next() {
		var table string
		if err := res.Scan(&table); err != nil {
			return tables, err
		}

		tables = append(tables, table)
	}

	if res.Err() != nil {
		return tables, res.Err()
	}

	return tables, nil
}

func (ss *SqlStore) truncateTables() error {
	tables, err := ss.getTables()
	if err != nil {
		return err
	}
	for _, table := range tables {
		if _, err := ss.master().Exec("TRUNCATE TABLE " + table); err != nil {
			return err
		}
	}
	return nil
}

func (ss *SqlStore) loadStores() {
	ss.stores = &stores{
		system:  newSystemStore(ss),
		films:   newFilmStore(ss),
		ratings: newRatingStore(ss),
	}
}

func (ss *SqlStore) Close() error {
	if err := ss.masterDB.Close(); err != nil {
		return err
	}
	return nil
}

func (ss *SqlStore) master() *sql.DB {
	return ss.masterDB
}

func (ss *SqlStore) replica() *sql.DB {
	if ss.replicaDBs == nil {
		return ss.masterDB
	}

	defer func() {
		if ss.replicaId == len(ss.replicaDBs)-1 {
			ss.replicaId = 0
			return
		}
		ss.replicaId++
	}()
	return ss.replicaDBs[ss.replicaId]
}

func (ss *SqlStore) System() store.SystemStore {
	return ss.stores.system
}

func (ss *SqlStore) Film() store.FilmStore {
	return ss.stores.films
}

func (ss *SqlStore) Ratings() store.RatingStore {
	return ss.stores.ratings
}

func (ss *SqlStore) Insert(ctx context.Context, query string, args ...interface{}) error {
	stmt, err := ss.master().PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		return err
	}

	return err
}
