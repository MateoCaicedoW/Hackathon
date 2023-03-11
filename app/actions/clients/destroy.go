package clients

import (
	"fmt"
	"hackathon/app/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
)

func Delete(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	client := models.Client{}

	if err := tx.Find(&client, uuid.FromStringOrNil(c.Param("client_id"))); err != nil {
		return err
	}

	fmt.Println("client ------------>", client)

	if err := tx.Destroy(&client); err != nil {
		return err
	}

	return c.Redirect(http.StatusFound, "/admin/clients")
}
