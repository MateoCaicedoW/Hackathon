package loan

import (
	"hackathon/app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func Edit(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	loan := &models.Loan{}
	clients := &[]models.Client{}

	if err := tx.Eager("Person").All(clients); err != nil {
		return err
	}

	if err := tx.Find(loan, c.Param("loan_id")); err != nil {
		return err
	}

	c.Set("clients", clients)
	c.Set("loan", loan)
	return c.Render(200, r.HTML("/loans/edit.plush.html"))

}

func Update(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	loan := models.Loan{}

	err := tx.Find(&loan, c.Param("loan_id"))
	if err != nil {
		return err
	}

	if err := c.Bind(&loan); err != nil {
		return err
	}

	err = tx.Eager("Client").Update(&loan)
	if err != nil {
		return err
	}
	return c.Redirect(302, "/admin/loans")
}
