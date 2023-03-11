CREATE TABLE loan_payments (
    id UUID PRIMARY KEY,
    mount NUMERIC NOT NULL,
    date TIMESTAMPTZ NOT NULL,
    loan_id UUID NOT NULL,
    CONSTRAINT fk_loan
        FOREIGN KEY(loan_id) 
        REFERENCES loans(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
)