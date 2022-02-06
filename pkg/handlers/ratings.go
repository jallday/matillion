package handlers

import (
	"net/http"

	"gitlab.com/joshuaAllday/matillion/pkg/utils/models"
)

func (handlers *Handlers) initRatings() {
	handlers.Routes.Ratings.Handle("", handlers.ApiHandler(createRating)).Methods("POST")
}

// swagger:operation POST /ratings Rating Rating
// ---
// summary: Create Rating
// parameters:
// - name: rating
//   in: body
//   schema:
//     "$ref": "#/definitions/Rating"
//   required: true
// responses:
//   '200':
//     description: rating response
//     schema:
//        "$ref": "#/definitions/Rating"
//   default:
//     description: error response
//     schema:
//       "$ref": "#/definitions/Error"
func createRating(c *Context, w http.ResponseWriter, req *http.Request) {
	rating, err := models.RatingFromJSON(req.Body)
	if err != nil {
		c.SetInvalidBody("rating")
		return
	}

	// Checking if the film exists in our db they are rating
	if _, nErr := c.App.GetFilm(c.Ctx, rating.FilmID); nErr != nil {
		c.Err = nErr
		return
	}

	rating, nErr := c.App.CreateRating(c.Ctx, rating)
	if nErr != nil {
		c.Err = nErr
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(rating.ToJSON()))
}
