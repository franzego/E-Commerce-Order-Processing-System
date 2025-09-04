-- name: CreateOrderByProducName :one
INSERT INTO orders (
    product_id, quantity, total, status
)
SELECT 
    i.product_id,
    $2::int,
    i.price * $2::numeric,
    'pending'
FROM inventory i
WHERE i.product_name = $1
RETURNING *; 

-- name: GetOrder :one
SELECT * FROM orders WHERE id = $1 LIMIT 1;

-- name: UpdateOrder :one
UPDATE orders SET status = $1 WHERE id = $2 RETURNING *;

-- name: ListOrders :many
SELECT * FROM orders ORDER BY created_at DESC;

-- name: DeleteOrder :exec
DELETE FROM orders WHERE id = $1;
