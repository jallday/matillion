package app

import (
	"context"
	"net/http"

	"gitlab.com/joshuaAllday/matillion/pkg/utils/models"
)

func (a *app) SystemHealthCheck(ctx context.Context) *models.Error {
	if err := a.Srv().Store.System().Ping(ctx); err != nil {
		return models.NewError("app.SystemHealthCheck", "app.system.health.check", err.Error(), nil, http.StatusInternalServerError)
	}
	return nil
}
