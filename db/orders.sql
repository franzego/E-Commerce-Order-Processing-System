-- name: CreateOrder :one
INSERT INTO orders (product_id, quantity, status, created_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetOrder :one
SELECT * FROM orders WHERE id = $1 LIMIT 1;

-- name: UpdateOrder :one
UPDATE orders SET status = $1 WHERE id = $2 RETURNING *;

-- name: ListOrders :many
SELECT * FROM orders ORDER BY created_at DESC;

-- name: DeleteOrder :exec
DELETE FROM orders WHERE id = $1;
