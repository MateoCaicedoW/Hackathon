package app

import (
	"net/http"

	"hackathon/app/actions/account"
	"hackathon/app/actions/auth"
	"hackathon/app/actions/clients"
	"hackathon/app/actions/home"
	"hackathon/app/actions/index"
	ls "hackathon/app/actions/loan"
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

	app.POST("/send-request", home.Create)

	authG := app.Group("/auth")
	authG.GET("/login", auth.AuthNew)
	authG.POST("/login", auth.AuthCreate).Name("signinPath")

	admin := app.Group("/admin")
	admin.GET("/", index.PageIndex)

	accounts := admin.Group("/accounts")
	accounts.GET("/", account.List)
	accounts.GET("/new", account.New)
	accounts.POST("/", account.Create)
	accounts.GET("/{id}", account.Show)
	accounts.GET("/{id}/edit", account.Edit)
	accounts.PUT("/{id}", account.Update).Name("updateAccountPath")
	accounts.DELETE("/{id}", account.Delete).Name("deleteAccountPath")
	app.POST("/send-request", home.Create)

	client := admin.Group("/clients")
	client.GET("/new", clients.New)
	client.POST("/new", clients.Create)
	client.GET("/", clients.List)

	loan := admin.Group("/loans")
	loan.GET("/", ls.List)
	loan.GET("/new", ls.New)

	app.ServeFiles("/", http.FS(public.FS()))

}
