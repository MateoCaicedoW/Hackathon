package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Person struct {
	ID        string `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Nit       string `json:"nit" db:"nit"`
	Email     string `json:"email" db:"email"`
	Phone     string `json:"phone" db:"phone"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type People []Person

// String is not required by pop and may be deleted

func (p Person) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

func (p *Person) Validate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.EmailIsPresent{Field: p.Email, Name: "Email"},
		&validators.StringIsPresent{Field: p.Nit, Message: "Password incorrect.", Name: "Nit"},
	), err
}
