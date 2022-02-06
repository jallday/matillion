package storetesting

import (
	"database/sql"
	"net/url"
	"path"
	"strings"

	"github.com/pkg/errors"
	"gitlab.com/joshuaAllday/matillion/pkg/utils"
	"gitlab.com/joshuaAllday/matillion/pkg/utils/models"
)

func dbId() string {
	return strings.ReplaceAll(utils.NewID(), "-", "")
}

func SetupTestingStore() (*models.SqlSettings, error) {
	// Build default settings
	settings := &models.SqlSettings{}
	settings.SetDefaults()

	// creating a new db
	dns, err := url.Parse(settings.MasterURL)
	if err != nil {
		return nil, err
	}
	dns.Path = "db" + dbId()
	settings.MasterURL = dns.String()

	// create the db and give all permissions to test user
	if err := rootExec(settings, "CREATE DATABASE "+dns.Path); err != nil {
		return nil, err
	}

	if err := rootExec(settings, "GRANT ALL PRIVILEGES ON DATABASE \""+dns.Path+"\" TO mtuser"); err != nil {
		return nil, err
	}

	return settings, nil
}

func CleanupTestingStore(settings *models.SqlSettings) error {

	dbName, err := dbNameFromPath(settings.MasterURL)
	if err != nil {
		return err
	}
	if err := rootExec(settings, "DROP DATABASE "+dbName); err != nil {
		return err
	}
	return nil
}

func dbNameFromPath(dsn string) (string, error) {
	dsnURL, err := url.Parse(dsn)
	if err != nil {
		return "", err
	}

	return path.Base(dsnURL.Path), nil
}

func rootFromDSN(dsn string) (string, error) {
	dsnURL, err := url.Parse(dsn)
	if err != nil {
		return "", err
	}
	dsnURL.Path = "postgres"
	return dsnURL.String(), nil

}

// execAsRoot executes the given sql as root against the testing database
func rootExec(settings *models.SqlSettings, sqlCommand string) error {
	dsn, err := rootFromDSN(settings.MasterURL)
	if err != nil {
		return err
	}
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return errors.Wrapf(err, "failed to connect to database as root")
	}
	defer db.Close()
	if _, err = db.Exec(sqlCommand); err != nil {
		return errors.Wrapf(err, "failed to execute `%s` against database as root", sqlCommand)
	}

	return nil
}
