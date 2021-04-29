BEGIN;
CREATE TABLE IF NOT EXISTS payment_code (
  id uuid NOT NULL PRIMARY KEY,
  payment_code TEXT NOT NULL,
  name TEXT NOT NULL,
  status TEXT NOT NULL,
  expiration_date timestamptz NOT NULL,
  created_at timestamptz NOT NULL,
  updated_at timestamptz NOT NULL
);
COMMIT;