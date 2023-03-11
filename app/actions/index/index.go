package index

import (
	"hackathon/app/render"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

var r = render.Engine

func PageIndex(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("/dashboard/index.plush.html"))
}
