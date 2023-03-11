package auth

import (
	"hackathon/app/render"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

var (
	r = render.Engine
)

func Login(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("auth/login.html"))
}

func Destroy(c buffalo.Context) error {
	c.Session().Clear()
	return c.Redirect(http.StatusFound, "/")
}
