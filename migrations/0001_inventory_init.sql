-- baseline: inventory-service owns the inventory table (R2 realistic-stand).
CREATE TABLE IF NOT EXISTS inventory (
  id         text PRIMARY KEY,
  sku        text NOT NULL,
  qty        integer NOT NULL,
  created_at timestamptz NOT NULL DEFAULT now()
);
