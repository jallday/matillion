package sqlstore

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/lib/pq"
	"gitlab.com/joshuaAllday/matillion/pkg/utils/models"
)

var (
	PingRetries int = 5
)

var (
	ErrConnection = errors.New("unable to connect to db")
)

type Flag string

var (
	MasterFlag  Flag = "Master"
	ReplicaFlag Flag = "Replica"
)

func connect(connectionName string, settings *models.SqlSettings, flag Flag) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionName)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	for i := 0; i < PingRetries; i++ {
		if err := db.PingContext(ctx); err == nil {
			break
		}

		if i == PingRetries-1 {
			return nil, ErrConnection
		}

		time.Sleep(1 * time.Second)
	}

	// TODO(Josh): connection details based off of the flag - master or replica

	return db, nil
}

func IsConflictError(err error, ids []string) bool {
	unique := false
	if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
		unique = true
	}

	field := false
	for _, id := range ids {
		if strings.Contains(err.Error(), id) {
			field = true
			break
		}
	}
	return unique && field
}
