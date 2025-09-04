-- Inventory table
CREATE TABLE inventory (
  product_id serial PRIMARY KEY,
  price numeric(12,2) NOT NULL CHECK (price >= 0),
  currency varchar(4) NOT NULL,
  stock int NOT NULL CHECK (stock >= 0),
  updated_at timestamptz DEFAULT now()
);

-- Orders table
CREATE TABLE orders (
  id serial PRIMARY KEY,
  product_id int NOT NULL REFERENCES inventory(product_id),
  quantity int NOT NULL CHECK (quantity > 0),
  total numeric(12,2),
  status varchar(20) NOT NULL DEFAULT 'pending',
  created_at timestamptz DEFAULT now()
);

