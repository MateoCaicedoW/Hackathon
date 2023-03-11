package app

import (
	"hackathon/internal/environment"

	"github.com/gobuffalo/buffalo"
)

// App creates a new application with default settings and reading
// GO_ENV. It calls setRoutes to setup the routes for the app that's being
// created before returning it
func New() (*buffalo.App, error) {
	app := buffalo.New(buffalo.Options{
		Env:         environment.Current(),
		SessionName: environment.SessionName,
		WorkerOff:   true,
	})

	setRoutes(app)

	return app, nil

}
