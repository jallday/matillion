package app

import (
	"context"
	"errors"
	"net/http"

	"gitlab.com/joshuaAllday/matillion/pkg/utils/models"
)

func (a *app) ListFilms(ctx context.Context, options *models.FilmOptions) ([]*models.Film, *models.Error) {
	films, err := a.Srv().Store.Film().List(ctx, options)
	if err != nil {
		return nil, models.NewError("app.ListFilms", "app.films.list.error", err.Error(), map[string]interface{}{
			"per_page": options.PerPage,
			"page":     options.Page,
		}, http.StatusInternalServerError)
	}
	return films, nil
}

func (a *app) GetFilm(ctx context.Context, id string) (*models.Film, *models.Error) {
	film, err := a.Srv().Store.Film().Get(ctx, id)
	if err != nil {
		var nErf *models.ErrNotFound
		switch {
		case errors.As(err, &nErf):
			return nil, models.NewError("app.GetFilm", "app.film.get.not_found", err.Error(), map[string]interface{}{
				"film_id": id,
			}, http.StatusNotFound)
		default:
			return nil, models.NewError("app.GetFilm", "app.film.get.error", err.Error(), nil, http.StatusInternalServerError)
		}
	}
	return film, nil
}
