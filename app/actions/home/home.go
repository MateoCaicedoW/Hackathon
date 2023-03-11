package home

import (
	"hackathon/app/models"
	"hackathon/app/render"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

var (
	// r is a buffalo/render Engine that will be used by actions
	// on this package to render render HTML or any other formats.
	r = render.Engine
)

func Index(c buffalo.Context) error {
	person := models.Person{}

	c.Set("person", person)
	return c.Render(http.StatusOK, r.HTML("home/index.plush.html"))
}

func Create(c buffalo.Context) error {
	person := models.Person{}
	tx := c.Value("tx").(*pop.Connection)
	if err := c.Bind(&person); err != nil {
		return err
	}

	verrs, err := person.Validate(tx)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		c.Set("person", person)
		c.Set("errors", verrs)
		return c.Render(http.StatusOK, r.HTML("home/index.plush.html"))
	}

	err = tx.Create(&person)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/")
}
