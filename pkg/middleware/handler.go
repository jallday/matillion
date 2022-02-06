package middleware

import (
	"context"
	"net/http"
	"time"

	"gitlab.com/joshuaAllday/matillion/pkg/app"
	"gitlab.com/joshuaAllday/matillion/pkg/utils/mtlog"
	responsewriter "gitlab.com/joshuaAllday/matillion/pkg/utils/response_writer"
)

type Handler struct {
	App            app.App
	Function       func(c *RequestContext, w http.ResponseWriter, req *http.Request)
	RequireSession bool
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	startTime := time.Now()
	sw := responsewriter.ResponseWriter{ResponseWriter: w}
	c := &RequestContext{
		Ctx: context.Background(),
		App: h.App,
	}
	c.loadRequestDetails(req)
	w.Header().Add("Content-type", "application/json")
	w.Header().Add("X-request-id", c.RequestID)

	h.App.Srv().Logger().Info("Starting http request", mtlog.Fields{
		"request_id": c.RequestID,
		"path":       req.URL.Path,
	})

	if h.RequireSession {
		// TODO(Josh): This is a great place to add any auth functionality whether it be
		// calling a function to our db or another service to validate any tokens etc
	}

	h.Function(c, &sw, req)

	if c.Err != nil {
		c.Err.RequestId = c.RequestID
		if c.Err.StatusCode == http.StatusInternalServerError {
			c.Err.DetailedError = "Internal Server Error"
			c.Err.ID = ""
			c.Err.Params = nil
		}
		sw.WriteHeader(c.Err.StatusCode)
		sw.Write([]byte(c.Err.ToJSON()))
	}

	h.App.Srv().Logger().Info("Finished http request", mtlog.Fields{
		"request_id":  c.RequestID,
		"elapsed":     float64(time.Since(startTime)) / float64(time.Second),
		"status_code": sw.StatusCode(),
	})

}
