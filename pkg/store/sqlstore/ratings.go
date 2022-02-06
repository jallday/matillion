package sqlstore

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"gitlab.com/joshuaAllday/matillion/pkg/utils/models"
)

type ratingsStore struct {
	*SqlStore
}

func newRatingStore(ss *SqlStore) *ratingsStore {
	return &ratingsStore{ss}
}

func (rs *ratingsStore) Save(ctx context.Context, rating *models.Rating) (*models.Rating, error) {
	rating.SaveSanitisation()
	if nErr := rating.IsValid(); nErr != nil {
		return nil, nErr
	}

	if err := rs.Insert(ctx, querySaveRating,
		rating.ID, rating.Author, rating.FilmID, rating.Score,
		rating.CreatedAt, rating.UpdatedAt, rating.DeletedAt,
	); err != nil {
		if IsConflictError(err, []string{"films_author_filmid_key"}) {
			return nil, &models.ErrConflict{Resource: "ratings", Param: fmt.Sprintf("%s.%s", rating.Author, rating.FilmID)}
		}

		return nil, errors.Wrapf(err, "unable to create rating for film=%v", rating.FilmID)
	}

	return rating, nil
}

func (rs *ratingsStore) List(ctx context.Context, options *models.RatingOptions) ([]*models.Rating, error) {
	rows, err := rs.replica().QueryContext(ctx, queryListRatings,
		options.Id, options.MinScore, options.MaxScore,
		options.PerPage, options.PerPage*options.Page)
	if err != nil {
		return nil, errors.Wrap(err, "unable to list ratings")
	}
	defer rows.Close()
	ratings := make([]*models.Rating, 0)
	for rows.Next() {
		rating := new(models.Rating)
		if err := rows.Scan(
			&rating.ID, &rating.Author, &rating.FilmID, &rating.Score,
			&rating.CreatedAt, &rating.UpdatedAt, &rating.DeletedAt,
		); err != nil {
			return nil, errors.Wrap(err, "problem with scan rating into ratings")
		}
		ratings = append(ratings, rating)
	}
	return ratings, nil
}

const (
	querySaveRating = `
		INSERT INTO ratings(
			id, author, film_id, score, created_at,
			updated_at, deleted_at
		) VALUES($1,$2,$3,$4,$5,$6,$7)
	`
	queryListRatings = `
		SELECT
			id, author, film_id, score, created_at,
			updated_at, deleted_at
		FROM ratings
		WHERE (
			Film_Id = $1 AND (score >= $2 AND score <= $3)
		)
		LIMIT $4
		OFFSET $5
	`
)
