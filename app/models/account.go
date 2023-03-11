package models

import "github.com/gofrs/uuid"

type Account struct {
	ID     uuid.UUID `json:"id" db:"id"`
	Name   string    `json:"name" db:"name"`
	Number string    `json:"number" db:"number"`
	Kind   string    `json:"kind" db:"kind"`
}
