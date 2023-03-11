package account

import (
	"hackathon/app/models"
	"hackathon/app/render"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

var r = render.Engine

func List(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	accounts := &[]models.Account{}

	if err := tx.All(accounts); err != nil {
		return err
	}

	c.Set("accounts", accounts)
	return c.Render(http.StatusOK, r.HTML("/accounts/list.plush.html"))
}
