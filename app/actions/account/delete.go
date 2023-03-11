package account

import (
	"hackathon/app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func Delete(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	account := models.Account{}

	if err := tx.Find(&account, c.Param("id")); err != nil {
		return err
	}

	if err := tx.Destroy(&account); err != nil {
		return err
	}

	return c.Redirect(302, "/admin/accounts")
}
