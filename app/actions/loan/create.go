package loan

import (
	"hackathon/app/models"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func New(c buffalo.Context) error {
	loan := &models.Loan{}

	clients := &[]models.Client{}
	tx := c.Value("tx").(*pop.Connection)
	if err := tx.Eager("Person").All(clients); err != nil {
		return err
	}
	loan.Date = time.Now()

	c.Set("clients", clients)

	c.Set("loan", loan)
	return c.Render(200, r.HTML("loans/new.html"))

}
func Create(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	loan := &models.Loan{}
	if err := c.Bind(loan); err != nil {
		return err
	}

	err := tx.Create(loan)
	if err != nil {
		return err
	}
	return c.Redirect(302, "/admin/loans")
}
