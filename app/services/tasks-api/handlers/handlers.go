package handlers

import (
	"net/http"
	"os"

	"github.com/Bruno-10/tasks/app/services/tasks-api/handlers/v1/taskgrp"
	"github.com/Bruno-10/tasks/business/core/task"
	"github.com/Bruno-10/tasks/business/core/task/stores/taskdb"
	"github.com/Bruno-10/tasks/business/web/v1/mid"
	"github.com/Bruno-10/tasks/foundation/web"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// APIMuxConfig contains all the mandatory systems required by handlers.
type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *zap.SugaredLogger
	DB       *sqlx.DB
}

// APIMux constructs a http.Handler with all application routes defined.
func APIMux(cfg APIMuxConfig) *web.App {
	app := web.NewApp(cfg.Shutdown, mid.Logger(cfg.Log), mid.Errors(cfg.Log), mid.Metrics(), mid.Panics())

	// -------------------------------------------------------------------------

	tskCore := task.NewCore(taskdb.NewStore(cfg.Log, cfg.DB))

	tgh := taskgrp.New(tskCore)

	app.Handle(http.MethodGet, "/tasks", tgh.Query)

	return app
}
