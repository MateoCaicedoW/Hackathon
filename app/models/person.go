package models

import (
	"regexp"
	"strings"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
)

type Person struct {
	ID        uuid.UUID `json:"id" db:"id"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	NIT       string    `json:"nit" db:"nit"`
	Email     string    `json:"email" db:"email"`
	Phone     string    `json:"phone" db:"phone"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (u *Person) Validate(tx *pop.Connection) (*validate.Errors, error) {

	return validate.Validate(
		&validators.StringIsPresent{Field: u.FirstName, Name: "First Name"},
		&validators.FuncValidator{
			Fn: func() bool {
				if u.FirstName != "" && !regexp.MustCompile(`^[a-zA-Z ]+$`).MatchString(u.FirstName) {
					return false
				}
				return true
			},
			Name:    "First Name",
			Message: "%s First Name must be letters only.",
		},
		&validators.FuncValidator{
			Fn: func() bool {
				if u.LastName != "" && !regexp.MustCompile(`^[a-zA-Z ]+$`).MatchString(u.LastName) {
					return false
				}
				return true
			},
			Name:    "Last Name",
			Message: "%s Last Name must be letters only.",
		},
		&validators.StringIsPresent{Field: u.LastName, Name: "Last Name"},
		&validators.StringIsPresent{Field: u.Email, Name: "Email"},
		&validators.StringIsPresent{
			Field: strings.ToLower(u.Phone),
			Name:  "Phone",
		},
		&validators.StringIsPresent{
			Field: strings.ToLower(u.NIT),
			Name:  "Nit",
		},

		&validators.FuncValidator{

			Fn: func() bool {
				if u.FirstName != "" && len(u.FirstName) > 50 && regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(u.FirstName) {
					return false
				}
				return true
			},
			Name:    "First Name",
			Message: "%s First Name must be less than 50 characters.",
		},
		&validators.FuncValidator{
			Fn: func() bool {
				if u.LastName != "" && len(u.LastName) > 50 && regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(u.LastName) {
					return false
				}
				return true
			},
			Name:    "Last Name",
			Message: "%s Last Name must be less than 50 characters.",
		},

		&validators.FuncValidator{

			Name:    "Email",
			Message: "%s Email is already taken",
			Fn: func() bool {
				var b bool
				q := tx.Where("email = ?", u.Email)
				if u.ID != uuid.Nil {
					q = q.Where("id != ?", u.ID)
				}
				b, err := q.Exists(u)
				if err != nil {
					return false
				}
				return !b
			},
		},
		&validators.FuncValidator{
			Fn: func() bool {
				ex := `^[\w\.]+@([\w-]+\.)+[\w-]{2,4}$`
				if u.Email != "" && !regexp.MustCompile(ex).MatchString(u.Email) {
					return false
				}
				return true
			},
			Name:    "Email",
			Message: "%s Email is invalid",
		},
		&validators.FuncValidator{
			Fn: func() bool {
				if u.Email != "" {
					local := strings.Split(u.Email, "@")
					str := local[0]
					if len(str) > 64 {
						return false
					}
				}
				return true
			},
			Name:    "Email",
			Message: "%s Before @ Email must be less or equal than 64 characters ",
		},
		&validators.FuncValidator{
			Fn: func() bool {
				if u.Email != "" {
					local := strings.Split(u.Email, "@")
					str := local[1]
					if len(str) > 255 {
						return false
					}
				}
				return true
			},
			Name:    "Email",
			Message: "%s After @ Email must be less or equal than 255 characters ",
		},

		&validators.FuncValidator{
			Fn: func() bool {
				if u.NIT != "" {
					if len(u.NIT) < 10 {
						return false
					}
				}
				return true
			},
			Name:    "Nit",
			Message: "%s NIT must be 10 characters.",
		},
		&validators.FuncValidator{
			Fn: func() bool {
				if u.Phone != "" && len(u.Phone) > 10 {
					return false
				}
				return true
			},
			Name:    "Phone",
			Message: "%s Phone must be less than 10 characters.",
		},
		&validators.FuncValidator{
			Fn: func() bool {
				if u.NIT != "" {
					if len(u.NIT) > 10 {
						return false
					}
				}

				return true
			},
			Name:    "Nit",
			Message: "%s Nit must be less than 10 characters.",
		},
		&validators.FuncValidator{
			Fn: func() bool {
				if u.Phone != "" && len(u.Phone) > 10 {
					return false
				}
				return true
			},
			Name:    "Phone",
			Message: "%s Phone must be less than 10 characters.",
		},

		&validators.FuncValidator{
			Fn: func() bool {
				if u.Phone != "" && !regexp.MustCompile(`^[0-9]+$`).MatchString(u.Phone) {
					return false
				}
				return true
			},
			Name:    "Phone",
			Message: "%s Phone must be numbers only.",
		},

		&validators.FuncValidator{
			Fn: func() bool {
				if u.NIT != "" {
					if regexp.MustCompile(`^[0-9]+$`).MatchString(u.NIT) {
						return true
					}
				}
				return false
			},
			Name:    "Nit",
			Message: "%s NIT must be numbers only.",
		},
	), nil
}
