package app

import (
	"context"

	"gitlab.com/joshuaAllday/matillion/pkg/server"
	"gitlab.com/joshuaAllday/matillion/pkg/utils/models"
)

type app struct {
	srv *server.Server
}

type App interface {
	Srv() *server.Server
	SetServer(s *server.Server)
	SystemHealthCheck(ctx context.Context) *models.Error
	ListFilms(ctx context.Context, options *models.FilmOptions) ([]*models.Film, *models.Error)
	GetFilm(ctx context.Context, id string) (*models.Film, *models.Error)
	CreateRating(ctx context.Context, rating *models.Rating) (*models.Rating, *models.Error)
	GetRatingsByFilm(ctx context.Context, options *models.RatingOptions) ([]*models.Rating, *models.Error)
}

func New(s *server.Server) App {
	a := &app{
		srv: s,
	}

	return a
}

func (a *app) Srv() *server.Server {
	return a.srv
}

func (a *app) SetServer(s *server.Server) {
	a.srv = s
}
