package handlers

import (
	"fmt"

	"github.com/gorilla/mux"
	"gitlab.com/joshuaAllday/matillion/pkg/app"
	"gitlab.com/joshuaAllday/matillion/pkg/utils/models"
)

type Handlers struct {
	Routes *Routes
	app    app.App
}

type Routes struct {
	Root    *mux.Router // ''
	ApiRoot *mux.Router // 'api/v1'

	Films   *mux.Router // 'api/v1/films'
	System  *mux.Router // 'api/v1/system'
	Ratings *mux.Router // 'api/v1/ratings

}

func InitHandlers(router *mux.Router, a app.App) *Handlers {
	handlers := &Handlers{
		Routes: &Routes{},
		app:    a,
	}

	handlers.Routes.Root = router
	handlers.Routes.ApiRoot = handlers.Routes.Root.PathPrefix(fmt.Sprintf("/api/%s", models.API_VERSION)).Subrouter()
	handlers.Routes.Films = handlers.Routes.ApiRoot.PathPrefix("/films").Subrouter()
	handlers.Routes.Ratings = handlers.Routes.ApiRoot.PathPrefix("/ratings").Subrouter()
	handlers.Routes.System = handlers.Routes.ApiRoot.PathPrefix("/system").Subrouter()

	handlers.initSystem()
	handlers.initFilms()
	handlers.initRatings()

	return handlers
}
