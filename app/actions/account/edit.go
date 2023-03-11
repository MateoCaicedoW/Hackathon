package account

import (
	"hackathon/app/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func Edit(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	account := models.Account{}

	if err := tx.Find(&account, c.Param("id")); err != nil {
		return err
	}

	c.Set("account", account)

	return c.Render(http.StatusOK, r.HTML("/accounts/edit.plush.html"))
}

func Update(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	account := models.Account{}

	if err := tx.Find(&account, c.Param("id")); err != nil {
		return err
	}

	if err := c.Bind(&account); err != nil {
		return err
	}

	err := tx.Update(&account)
	if err != nil {
		return err
	}

	c.Set("account", account)
	return c.Redirect(http.StatusSeeOther, "/admin/accounts")
}
