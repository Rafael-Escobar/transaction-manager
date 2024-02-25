BEGIN;

-- Create the event_type table
CREATE TABLE IF NOT EXISTS transactions (
    id BIGSERIAL PRIMARY KEY,
    account_id BIGINT NOT NULL,
    event_type_id BIGINT NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    event_date TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES accounts(id),
    FOREIGN KEY (event_type_id) REFERENCES operation_types(id)
);

COMMIT;