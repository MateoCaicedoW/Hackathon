package models

import "github.com/gofrs/uuid"

type PaymentDetail struct {
	ID        uuid.UUID `json:"id" db:"id"`
	AccountID uuid.UUID `json:"account_id" db:"account_id"`
	Details   string    `json:"details" db:"details"`

	Account Account `json:"account" db:"-" belongs_to:"account" fk_id:"account_id"`
}
