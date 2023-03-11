package app

import (
	"net/http"

	"hackathon/app/actions/auth"
	"hackathon/app/actions/home"
	"hackathon/app/middleware"
	"hackathon/public"

	"github.com/gobuffalo/buffalo"
)

// SetRoutes for the application
func setRoutes(app *buffalo.App) {
	app.Use(middleware.RequestID)
	app.Use(middleware.Database)
	app.Use(middleware.ParameterLogger)
	app.Use(middleware.CSRF)

	app.GET("/", home.Index)
	authG := app.Group("/auth")
	authG.GET("/login", auth.Login)

	app.POST("/send-request", home.Create)

	app.ServeFiles("/", http.FS(public.FS()))

}
