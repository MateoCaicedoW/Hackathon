package clients

import (
	"hackathon/app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func List(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	clients := &models.Clients{}
	if err := tx.Eager().All(clients); err != nil {
		return err
	}

	c.Set("clients", clients)
	return c.Render(200, r.HTML("clients/list.plush.html"))
}
