package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Loan struct {
	ID         uuid.UUID `json:"id" db:"id"`
	Mount      float64   `json:"mount" db:"mount"`
	Commission float64   `json:"commission" db:"commission"`
	Date       time.Time `json:"date" db:"date"`
	ClientID   uuid.UUID `json:"client_id" db:"client_id"`
	Status     bool      `json:"status" db:"status"`

	Client Client `json:"client" db:"-" belongs_to:"client" fk_id:"client_id"`
}
