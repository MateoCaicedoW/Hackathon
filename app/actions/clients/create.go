package clients

import (
	"hackathon/app/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func New(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	people := &[]models.Person{}
	client := &models.Client{}
	accounts := &[]models.Account{}

	if err := tx.All(people); err != nil {
		return err
	}

	if err := tx.All(accounts); err != nil {
		return err
	}

	c.Set("accounts", accounts)
	c.Set("people", people)
	c.Set("client", client)

	return c.Render(http.StatusOK, r.HTML("clients/new.plush.html"))
}

func Create(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	client := &models.Client{}
	if err := c.Bind(client); err != nil {
		return err
	}

	err := tx.Create(client)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/admin/clients")

}
