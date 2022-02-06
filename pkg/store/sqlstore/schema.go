package sqlstore

import (
	"gitlab.com/joshuaAllday/matillion/pkg/utils"
	"gitlab.com/joshuaAllday/matillion/scripts"
)

// loadSchema : loads the database models for this service
func (ss *SqlStore) loadSchema() error {
	schema, err := utils.ReadFile(scripts.Path("./db/postgres.sql"))
	if err != nil {
		return err
	}

	if _, err = ss.master().Exec(string(schema)); err != nil {
		return err
	}

	return nil
}

// loadSeeds : loads some test data into the database
func (ss *SqlStore) loadSeeds() error {
	seeds, err := utils.ReadFile(scripts.Path("./db/seed.sql"))
	if err != nil {
		return err
	}

	if _, err = ss.master().Exec(string(seeds)); err != nil {
		return err
	}

	return nil
}
