BEGIN;
CREATE TABLE IF NOT EXISTS inquiries (
  id uuid NOT NULL PRIMARY KEY,
  payment_code TEXT NOT NULL,
  transaction_id TEXT NOT NULL,
  created_at timestamptz NOT NULL,
  updated_at timestamptz NOT NULL
);
COMMIT;