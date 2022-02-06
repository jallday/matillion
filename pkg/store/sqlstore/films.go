package sqlstore

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"gitlab.com/joshuaAllday/matillion/pkg/utils/models"
)

type filmStore struct {
	*SqlStore
}

func newFilmStore(ss *SqlStore) *filmStore {
	return &filmStore{ss}
}

func (fs *filmStore) Get(ctx context.Context, id string) (*models.Film, error) {
	film := new(models.Film)
	row := fs.replica().QueryRowContext(ctx, queryGetFilm, id)
	if err := row.Scan(
		&film.ID, &film.Title, &film.EpisodeID, &film.Director,
		&film.Producer, &film.ReleaseDate, &film.CreatedAt,
		&film.UpdatedAt, &film.DeletedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, models.NewErrNotFound("film", id)
		}
		return nil, errors.Wrapf(err, "unable to get film with id=%s", id)
	}

	return film, nil
}

func (fs *filmStore) List(ctx context.Context, options *models.FilmOptions) ([]*models.Film, error) {
	rows, err := fs.replica().QueryContext(ctx, queryListFilms, options.PerPage, options.PerPage*options.Page)
	if err != nil {
		return nil, errors.Wrap(err, "unable to list films")
	}
	defer rows.Close()
	films := make([]*models.Film, 0)
	for rows.Next() {
		film := new(models.Film)
		if err := rows.Scan(
			&film.ID, &film.Title, &film.EpisodeID, &film.Director,
			&film.Producer, &film.ReleaseDate, &film.CreatedAt,
			&film.UpdatedAt, &film.DeletedAt,
		); err != nil {
			return nil, errors.Wrap(err, "problem with scan film into films")
		}
		films = append(films, film)
	}
	return films, nil
}

const (
	queryListFilms = `
		SELECT 
			id, title, episode_id, director, producer, 
			release_date, created_at, updated_at, deleted_at
		FROM films
		ORDER BY release_date
		LIMIT $1
		OFFSET $2
	`
	queryGetFilm = `
		SELECT 
			id, title, episode_id, director, producer, 
			release_date, created_at, updated_at, deleted_at
		FROM films
		WHERE id = $1
	`
)
