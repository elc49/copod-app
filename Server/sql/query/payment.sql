-- name: CreatePayment :one
INSERT INTO payments (
  email, amount, currency, reason, status, reference_id, title_id, wallet_address
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetPaymentByReferenceID :one
SELECT * FROM payments
WHERE reference_id = $1
LIMIT 1;

-- name: UpdatePaymentStatus :one
UPDATE payments SET status = $1
WHERE reference_id = $2
RETURNING *;

-- name: GetPaymentTitleByID :one
SELECT * FROM titles
WHERE id = $1;

-- name: GetPaymentsByStatus :many
SELECT * FROM payments
WHERE status = $1;
