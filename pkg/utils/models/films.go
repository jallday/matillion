package models

import (
	"encoding/json"
	"time"
)

// swagger:model
type Film struct {
	// Id of the film
	// readOnly: true
	ID string `json:"id"`
	// Title of the film
	// required: true
	// max length: 128
	// min length: 1
	Title string `json:"title"`
	// Episode number of the film
	// required: true
	EpisodeID int `json:"episode_id"`
	// The name of the director of the film
	// max length: 128
	Director string `json:"director"`
	// The name of the producer of the film
	// max length: 128
	Producer string `json:"producer"`
	// The time the film was released
	ReleaseDate time.Time `json:"release_date,omitempty"`
	// time the film was created at
	// readOnly: true
	CreatedAt int `json:"created_at"`
	// time the film was last updated at
	// readOnly: true
	UpdatedAt int `json:"updated_at"`
	// time film was deleted at
	// readOnly: true
	DeletedAt int `json:"deleted_at"`
}

type FilmOptions struct {
	Page    int
	PerPage int
}

// swagger:model
type FilmRatings struct {
	*Film
	// ratings of the film
	// readOnly: true
	Ratings []*Rating `json:"ratings"`
}

func (fr *FilmRatings) ToJSON() string {
	b, _ := json.Marshal(fr)
	return string(b)
}
