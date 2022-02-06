package handlers

import (
	"net/http"

	"gitlab.com/joshuaAllday/matillion/pkg/utils/models"
)

func (handlers *Handlers) initSystem() {
	handlers.Routes.System.Handle("/ping", handlers.ApiHandler(systemPing)).Methods("GET")
}

// swagger:operation GET /system/ping System System
// ---
// summary: Ping System
// responses:
//   '200':
//     description: system response
//     schema:
//        "$ref": "#/definitions/System"
//   default:
//     description: error response
//     schema:
//       "$ref": "#/definitions/Error"
func systemPing(c *Context, w http.ResponseWriter, req *http.Request) {
	if nErr := c.App.SystemHealthCheck(c.Ctx); nErr != nil {
		c.Err = nErr
		return
	}

	system := &models.System{
		Status:  "Healthy",
		Message: "The system is running within it's limits and is able to connect to the required services",
		Version: models.VERSION,
	}
	w.Write([]byte(system.ToJSON()))
}
