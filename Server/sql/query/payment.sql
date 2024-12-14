-- name: CreatePayment :one
INSERT INTO payments (
  email, amount, currency, reason, status, reference_id, onboarding_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetPaymentByReferenceID :one
SELECT * FROM payments
WHERE reference_id = $1
LIMIT 1;

-- name: UpdatePaymentStatus :one
UPDATE payments SET status = $1
WHERE reference_id = $2
RETURNING *;

-- name: GetPaymentOnboardingByID :one
SELECT * FROM onboardings
WHERE id = $1;

-- name: GetPaymentsByStatus :many
SELECT * FROM payments
WHERE status = $1;
