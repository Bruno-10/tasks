// Package v1 contains the full set of handler functions and routes
// supported by the v1 web api.
package v1

import (
	"net/http"

	"github.com/Bruno-10/tasks/app/services/tasks-api/handlers/v1/checkgrp"
	"github.com/Bruno-10/tasks/app/services/tasks-api/handlers/v1/taskgrp"
	"github.com/Bruno-10/tasks/business/core/task"
	"github.com/Bruno-10/tasks/business/core/task/stores/taskdb"
	"github.com/Bruno-10/tasks/foundation/web"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Config struct {
	Build string
	Log   *zap.SugaredLogger
	DB    *sqlx.DB
}

func Routes(app *web.App, cfg Config) {

	// -------------------------------------------------------------------------

	cgh := checkgrp.New(cfg.Build, cfg.DB)

	app.HandleNoMiddleware(http.MethodGet, "/readiness", cgh.Readiness)
	app.HandleNoMiddleware(http.MethodGet, "/liveness", cgh.Liveness)

	tskCore := task.NewCore(taskdb.NewStore(cfg.Log, cfg.DB))

	tgh := taskgrp.New(tskCore)

	app.Handle(http.MethodPost, "/tasks", tgh.Create)
	app.Handle(http.MethodGet, "/tasks", tgh.Query)
}
