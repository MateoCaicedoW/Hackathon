package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id" db:"id"`
	PersonID uuid.UUID `json:"person_id" db:"person_id"`
	Since    time.Time `json:"since" db:"since"`
	Until    time.Time `json:"until" db:"until"`

	Person Person `json:"person" db:"-" belongs_to:"person" fk_id:"person_id"`
}
