package models

import (
	"strings"
	"time"

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

func (p Person) SelectLabel() string {
	return strings.Join([]string{p.FirstName, p.LastName}, " ")
}

func (p Person) SelectValue() interface{} {
	return p.ID
}

func (p Person) FullName() string {
	return strings.Join([]string{p.FirstName, p.LastName}, " ")
}
