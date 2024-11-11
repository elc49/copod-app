-- name: CreatePayment :one
INSERT INTO payments (
  email, amount, currency, reason, status, reference_id, upload_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: ClearTestPayments :exec
DELETE FROM payments;
