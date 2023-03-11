CREATE TABLE clients (
    id UUID PRIMARY KEY,
    person_id UUID NOT NULL,
    account_id UUID NOT NULL,
    since TIMESTAMPTZ NOT NULL,
    until TIMESTAMPTZ NOT NULL,

    CONSTRAINT fk_person
        FOREIGN KEY(person_id) 
        REFERENCES people(id),
    
    CONSTRAINT fk_account
        FOREIGN KEY(account_id)
        REFERENCES accounts(id),

        
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
)