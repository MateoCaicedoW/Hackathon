package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Client struct {
	ID        uuid.UUID `json:"id" db:"id"`
	PersonID  uuid.UUID `json:"person_id" db:"person_id"`
	Since     time.Time `json:"since" db:"since"`
	Until     time.Time `json:"until" db:"until"`
	AccountID uuid.UUID `json:"account_id" db:"account_id"`

	Person  Person  `json:"person" db:"-" belongs_to:"person" fk_id:"person_id"`
	Account Account `json:"account" db:"-" belongs_to:"account" fk_id:"account_id"`
}
