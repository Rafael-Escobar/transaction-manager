BEGIN;

-- Drop indexes for the accounts table
DROP INDEX IF EXISTS idx_accounts_document_number;

DROP TABLE IF EXISTS accounts;

COMMIT;