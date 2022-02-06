package handlers

import (
	"encoding/json"
	"net/http"

	"gitlab.com/joshuaAllday/matillion/pkg/utils/models"
)

func (handlers *Handlers) initFilms() {
	handlers.Routes.Films.Handle("", handlers.ApiHandler(listFilms)).Methods("GET")
	handlers.Routes.Films.Handle("/{film_id}/ratings", handlers.ApiHandler(listFilmsRatings)).Methods("GET")

}

// swagger:operation GET /films Films Films
// ---
// summary: List Films
// parameters:
// - name: page
//   in: query
//   description: the page number
//   required: false
//   type: integer
//   format: int32
// - name: per_page
//   in: query
//   description: the number of results per page
//   required: false
//   type: integer
//   format: int32
// responses:
//   '200':
//     description: films response
//     schema:
//        items:
//           "$ref": "#/definitions/Film"
//   default:
//     description: error response
//     schema:
//       "$ref": "#/definitions/Error"
func listFilms(c *Context, w http.ResponseWriter, req *http.Request) {
	options := &models.FilmOptions{
		Page:    c.Params.Page,
		PerPage: c.Params.PerPage,
	}

	films, nErr := c.App.ListFilms(c.Ctx, options)
	if nErr != nil {
		c.Err = nErr
		return
	}

	b, err := json.Marshal(films)
	if err != nil {
		return
	}
	w.Write(b)
}

// swagger:operation GET /films/{film_id}/ratings Films Films
// ---
// summary: List Film's Rating
// parameters:
// - name: film id
//   in: path
//   required: true
// - name: page
//   in: query
//   description: the page number
//   required: false
//   type: integer
//   format: int32
// - name: per_page
//   in: query
//   description: the number of results per page
//   required: false
//   type: integer
//   format: int32
// - name: max_score
//   in: query
//   description: the maximum score for ratings
//   required: false
//   type: integer
//   format: int32
// - name: min_score
//   in: query
//   description: the minimum score for ratings
//   required: false
//   type: integer
//   format: int32
// responses:
//   '200':
//     description: film's rating response
//     schema:
//        "$ref": "#/definitions/FilmRatings"
//   default:
//     description: error response
//     schema:
//       "$ref": "#/definitions/Error"
func listFilmsRatings(c *Context, w http.ResponseWriter, req *http.Request) {
	options := &models.RatingOptions{
		Page:     c.Params.Page,
		PerPage:  c.Params.PerPage,
		Id:       c.Params.FilmID,
		MaxScore: c.Params.MaxScore,
		MinScore: c.Params.MinScore,
	}

	film, nErr := c.App.GetFilm(c.Ctx, options.Id)
	if nErr != nil {
		c.Err = nErr
		return
	}

	ratings, nErr := c.App.GetRatingsByFilm(c.Ctx, options)
	if nErr != nil {
		c.Err = nErr
		return
	}

	filmRatings := &models.FilmRatings{
		Film:    film,
		Ratings: ratings,
	}

	w.Write([]byte(filmRatings.ToJSON()))
}
