package middleware

import (
	"context"
	"net/http"

	"gitlab.com/joshuaAllday/matillion/pkg/app"
	"gitlab.com/joshuaAllday/matillion/pkg/utils"
	"gitlab.com/joshuaAllday/matillion/pkg/utils/models"
)

type RequestContext struct {
	Ctx context.Context
	RequestDetails
	App app.App
	Err *models.Error
}

type RequestDetails struct {
	RequestID string
	Path      string
	Params    *Params
}

func (c *RequestContext) loadRequestDetails(req *http.Request) {
	c.RequestDetails = RequestDetails{
		RequestID: utils.NewID(),
		Params:    LoadParams(req),
	}
}

func (c *RequestContext) SetInvalidBody(parameter string) {
	c.Err = models.NewError("middleware.SetInvalidBody", "api.invalid.body.error", "invalid body sent", map[string]interface{}{
		"parameter": parameter,
	}, http.StatusBadRequest)
}
