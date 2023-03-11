package account

import (
	"hackathon/app/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func New(c buffalo.Context) error {
	account := models.Account{}

	c.Set("account", account)
	return c.Render(http.StatusOK, r.HTML("/accounts/new.plush.html"))
}

func Create(c buffalo.Context) error {
	account := models.Account{}
	tx := c.Value("tx").(*pop.Connection)
	if err := c.Bind(&account); err != nil {
		return err
	}

	err := tx.Create(&account)
	if err != nil {
		return err
	}
	return c.Redirect(http.StatusSeeOther, "/admin/accounts")
}
