package handlers

import (
	"net/http"

	"gitlab.com/joshuaAllday/matillion/pkg/middleware"
)

type Context = middleware.RequestContext

// ApiHandler : this is used for apis which are meant to not require a session
// e.g system health check. Otherwise we should be using all the session or token ones.
func (h *Handlers) ApiHandler(f func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	return &middleware.Handler{
		App:            h.app,
		Function:       f,
		RequireSession: false,
	}
}

// When auth is setup use this one :)
func (h *Handlers) ApiSessionHandler(f func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	return &middleware.Handler{
		Function:       f,
		RequireSession: true,
	}
}
