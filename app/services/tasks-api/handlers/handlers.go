package handlers

import (
	"os"

	v1 "github.com/Bruno-10/tasks/app/services/tasks-api/handlers/v1"
	"github.com/Bruno-10/tasks/business/web/v1/mid"
	"github.com/Bruno-10/tasks/foundation/web"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// Options represent optional parameters.
type Options struct {
	corsOrigin string
}

// WithCORS provides configuration options for CORS.
func WithCORS(origin string) func(opts *Options) {
	return func(opts *Options) {
		opts.corsOrigin = origin
	}
}

// APIMuxConfig contains all the mandatory systems required by handlers.
type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *zap.SugaredLogger
	DB       *sqlx.DB
}

// APIMux constructs a http.Handler with all application routes defined.
func APIMux(cfg APIMuxConfig, options ...func(opts *Options)) *web.App {
	var opts Options
	for _, option := range options {
		option(&opts)
	}

	app := web.NewApp(cfg.Shutdown, mid.Logger(cfg.Log), mid.Errors(cfg.Log), mid.Metrics(), mid.Panics())

	if opts.corsOrigin != "" {
		app.EnableCORS(mid.Cors(opts.corsOrigin))
	}

	v1.Routes(app, v1.Config{
		Log: cfg.Log,
		DB:  cfg.DB,
	})

	return app
}
