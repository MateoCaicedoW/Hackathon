package clients

import (
	"hackathon/app/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func Edit(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	client := models.Client{}

	if err := tx.Find(&client, c.Param("client_id")); err != nil {
		return err
	}

	accounts := &[]models.Account{}
	people := &[]models.Person{}
	if err := tx.All(people); err != nil {
		return err
	}

	if err := tx.All(accounts); err != nil {
		return err
	}

	c.Set("accounts", accounts)
	c.Set("people", people)

	c.Set("client", client)
	return c.Render(http.StatusOK, r.HTML("/clients/edit.plush.html"))
}

func Update(c buffalo.Context) error {

	tx := c.Value("tx").(*pop.Connection)
	client := models.Client{}

	if err := tx.Find(&client, c.Param("client_id")); err != nil {
		return err
	}

	if err := c.Bind(&client); err != nil {
		return err
	}

	err := tx.Update(&client)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/admin/clients")

}
