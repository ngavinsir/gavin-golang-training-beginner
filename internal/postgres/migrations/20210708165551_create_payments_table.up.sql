BEGIN;
CREATE TABLE IF NOT EXISTS payments (
  id uuid NOT NULL PRIMARY KEY,
  payment_code TEXT NOT NULL,
  transaction_id TEXT NOT NULL,
  name TEXT NOT NULL,
  amount TEXT NOT NULL,
  created_at timestamptz NOT NULL,
  updated_at timestamptz NOT NULL
);
COMMIT;