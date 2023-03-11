package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type LoanPayment struct {
	ID     uuid.UUID `json:"id" db:"id"`
	LoanID uuid.UUID `json:"loan_id" db:"loan_id"`
	Mount  float64   `json:"mount" db:"mount"`
	Date   time.Time `json:"date" db:"date"`
}
