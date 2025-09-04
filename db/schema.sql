CREATE TABLE inventory (
  product_id serial PRIMARY KEY,
  price numeric(12,2) NOT NULL CHECK (price >= 0),
  currency varchar(4) NOT NULL,
  stock int NOT NULL CHECK (stock >= 0),
  updated_at timestamptz DEFAULT now()
);