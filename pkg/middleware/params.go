package middleware

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gitlab.com/joshuaAllday/matillion/pkg/utils/models"
)

var (
	DefaultPerPage = 60
	MaxPerPage     = 250
)

type Params struct {
	FilmID   string
	Email    string
	PerPage  int
	Page     int
	MaxScore int
	MinScore int
}

func LoadParams(req *http.Request) *Params {
	vars := mux.Vars(req)
	p := &Params{}
	if i, err := strconv.Atoi(req.URL.Query().Get("page")); err == nil {
		p.Page = i
	}

	if filmID, ok := vars["film_id"]; ok {
		p.FilmID = filmID
	}

	if i, err := strconv.Atoi(req.URL.Query().Get("max_score")); err == nil {
		p.MaxScore = i
	} else {
		p.MaxScore = models.MaxScore
	}

	if i, err := strconv.Atoi(req.URL.Query().Get("min_score")); err == nil {
		p.MinScore = i
	} else {
		p.MinScore = models.MinScore
	}

	p.calcPerPage(req)
	return p
}

func (p *Params) calcPerPage(req *http.Request) {
	i, err := strconv.Atoi(req.URL.Query().Get("per_page"))
	if err != nil || i < 0 {
		p.PerPage = DefaultPerPage
		return
	} else if i > MaxPerPage {
		p.PerPage = MaxPerPage
		return
	}
	p.PerPage = i
}
