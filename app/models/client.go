package models

import (
	"github.com/gofrs/uuid"
)

type Client struct {
	ID        uuid.UUID `json:"id" db:"id"`
	PersonID  uuid.UUID `json:"person_id" db:"person_id"`
	AccountID uuid.UUID `json:"account_id" db:"account_id"`

	Person  Person  `json:"person" db:"-" belongs_to:"person" fk_id:"person_id"`
	Account Account `json:"account" db:"-" belongs_to:"account" fk_id:"account_id"`
}

type Clients []Client

func (c Client) SelectLabel() string {
	return c.Person.FirstName + " " + c.Person.LastName
}

func (c Client) SelectValue() interface{} {
	return c.ID
}
