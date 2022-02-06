package models

import (
	"encoding/json"
	"io"
	"net/http"
	"unicode/utf8"

	"gitlab.com/joshuaAllday/matillion/pkg/utils"
)

var (
	MaxScore = 5
	MinScore = 0
)

// swagger:model
type Rating struct {
	// Id of the rating
	// readOnly: true
	ID string `json:"id"`
	// Author of the rating
	// required: true
	// max length: 128
	// min length: 1
	Author string `json:"author"`
	// Id of the film the rating relates too
	// required: true
	// pattern: [A-Za-z0-9_+]
	FilmID string `json:"film_id"`
	// The name of the score of the film
	// required: true
	// max: 5
	// min: 0
	Score int `json:"score"`
	// time the rating was created at
	// readOnly: true
	CreatedAt int `json:"created_at"`
	// time the rating was last updated at
	// readOnly: true
	UpdatedAt int `json:"updated_at"`
	// time rating was deleted at
	// readOnly: true
	DeletedAt int `json:"deleted_at"`
}

func RatingFromJSON(data io.Reader) (*Rating, error) {
	var r Rating
	if err := json.NewDecoder(data).Decode(&r); err != nil {
		return nil, err
	}
	return &r, nil
}

func (r *Rating) ToJSON() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *Rating) SaveSanitisation() {
	r.ID = utils.NewID()
	r.CreatedAt = utils.IntTime()
	r.UpdatedAt = r.CreatedAt
}

func (r *Rating) IsValid() *Error {
	if !utils.ValidId(r.ID) {
		return NewError("models.Rating.IsValid", "models.ratings.invalid.id", "invalid id", map[string]interface{}{"id": r.ID}, http.StatusBadRequest)
	}

	if r.CreatedAt == 0 {
		return NewError("models.Rating.IsValid", "models.ratings.invalid.created_at", "invalid created at time", nil, http.StatusBadRequest)
	}

	if r.UpdatedAt == 0 {
		return NewError("models.Rating.IsValid", "models.ratings.invalid.updated_at", "invalid updated at time", nil, http.StatusBadRequest)
	}

	if !utils.ValidId(r.FilmID) {
		return NewError("models.Rating.IsValid", "models.ratings.invalid.film_id", "invalid film id", map[string]interface{}{"film_id": r.FilmID}, http.StatusBadRequest)
	}

	if r.Score < MinScore || r.Score > MaxScore {
		return NewError("models.Rating.IsValid", "models.ratings.invalid.score", "invalid score - has to be a number between 0 and 5", map[string]interface{}{"score": r.Score}, http.StatusBadRequest)
	}

	if r.Author == "" || utf8.RuneCountInString(r.Author) > 128 {
		return NewError("models.Rating.IsValid", "models.ratings.invalid.author", "invalid author - has to be a string of length between 1 and 128 char", map[string]interface{}{"author": r.Author}, http.StatusBadRequest)
	}

	return nil
}

type RatingOptions struct {
	Id       string
	Page     int
	PerPage  int
	MaxScore int
	MinScore int
}
