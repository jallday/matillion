package app

import (
	"context"
	"errors"
	"net/http"

	"gitlab.com/joshuaAllday/matillion/pkg/utils/models"
)

func (a *app) CreateRating(ctx context.Context, rating *models.Rating) (*models.Rating, *models.Error) {
	nRating, err := a.Srv().Store.Ratings().Save(ctx, rating)
	if err != nil {
		var nErr *models.Error
		var nErc *models.ErrConflict
		switch {
		case errors.As(err, &nErr):
			return nil, nErr
		case errors.As(err, &nErc):
			return nil, models.NewError("app.CreateRating", "app.rating.save.conflict", "film already has been rated by this author", map[string]interface{}{
				"author":  rating.Author,
				"film_id": rating.FilmID,
			}, http.StatusConflict)
		default:
			return nil, models.NewError("app.CreateRating", "app.rating.save.error", err.Error(), nil, http.StatusInternalServerError)
		}
	}
	return nRating, nil
}

func (a *app) GetRatingsByFilm(ctx context.Context, options *models.RatingOptions) ([]*models.Rating, *models.Error) {
	ratings, err := a.Srv().Store.Ratings().List(ctx, options)
	if err != nil {
		return nil, models.NewError("app.GetRatingsByFilm", "app.ratings.list.error", err.Error(), map[string]interface{}{
			"per_page": options.PerPage,
			"page":     options.Page,
			"film_id":  options.Id,
		}, http.StatusInternalServerError)
	}
	return ratings, nil
}
