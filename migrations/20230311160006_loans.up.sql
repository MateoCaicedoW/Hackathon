CREATE TABLE loans (
    id UUID PRIMARY KEY,
    mount NUMERIC NOT NULL,
    commision NUMERIC NOT NULL,
    date TIMESTAMPTZ NOT NULL,
    status bool NOT NULL,
    client_id UUID NOT NULL,


    CONSTRAINT fk_client
        FOREIGN KEY(client_id) 
        REFERENCES clients(id),

        
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
)