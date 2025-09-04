-- name: GetInventory :one
SELECT * FROM inventory
WHERE product_id = $1 LIMIT 1;

-- name: ListInventory :many
SELECT * FROM inventory
ORDER BY product_id;

-- name: CreateInventory :one
INSERT INTO inventory (
    price, currency, stock, updated_at
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateInventory :exec
UPDATE inventory
  set price = $2,
  currency = $3,
  stock = $4,
  updated_at = $5
WHERE product_id = $1;

-- name: DeleteInventory :exec
DELETE FROM inventory
WHERE product_id = $1;