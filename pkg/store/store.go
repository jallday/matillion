package store

import (
	"context"

	"gitlab.com/joshuaAllday/matillion/pkg/utils/models"
)

type Store interface {
	System() SystemStore
	Film() FilmStore
	Ratings() RatingStore
	Close() error
}

type SystemStore interface {
	Ping(ctx context.Context) error
}

type FilmStore interface {
	Get(ctx context.Context, id string) (*models.Film, error)
	List(ctx context.Context, options *models.FilmOptions) ([]*models.Film, error)
}

type RatingStore interface {
	Save(ctx context.Context, rating *models.Rating) (*models.Rating, error)
	List(ctx context.Context, options *models.RatingOptions) ([]*models.Rating, error)
}
