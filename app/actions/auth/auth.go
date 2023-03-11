package auth

import (
	"fmt"
	"hackathon/app/models"
	"hackathon/app/render"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

var (
	r = render.Engine
)

func AuthNew(c buffalo.Context) error {
	c.Set("user", models.Person{})
	return c.Render(http.StatusOK, r.HTML("auth/login.html"))
}

func AuthCreate(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	u := &models.Person{}
	if err := c.Bind(u); err != nil {
		return fmt.Errorf("error binding user: %w", err)
	}

	verrs, err := u.Validate(tx)
	if err != nil {
		return fmt.Errorf("error validating user: %w", err)
	}

	if verrs.HasAny() {
		c.Set("user", u)
		c.Set("errors", verrs)
		return c.Render(http.StatusUnprocessableEntity, r.HTML("auth/login.html"))
	}

	if err := tx.Where("email = ? AND nit = ?", u.Email, u.NIT).First(u); err != nil {
		return fmt.Errorf("error finding user: %w", err)
	}

	c.Session().Set("current_user_id", u.ID)
	c.Flash().Add("success", "Welcome Back!")

	return c.Redirect(http.StatusFound, "/admin")
}

func Destroy(c buffalo.Context) error {
	c.Session().Clear()
	c.Flash().Add("success", "You've been logged out successfully!")
	return c.Redirect(http.StatusFound, "/auth/login")
}
