package account

import (
	"hackathon/app/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func Show(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	account := models.Account{}

	if err := tx.Find(&account, c.Param("id")); err != nil {
		return err
	}

	c.Set("account", account)
	return c.Render(http.StatusOK, r.HTML("/accounts/new.plush.html"))
}
