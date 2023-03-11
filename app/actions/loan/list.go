package loan

import (
	"hackathon/app/models"
	"hackathon/app/render"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

var (
	// r is a buffalo/render Engine that will be used by actions
	// on this package to render render HTML or any other formats.
	r = render.Engine
)

func List(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	loans := &[]models.Loan{}
	if err := tx.All(loans); err != nil {
		return err
	}

	c.Set("loans", loans)
	return c.Render(200, r.HTML("loans/list.html"))
}
