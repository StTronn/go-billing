-- name: CreateBillingMetric :execresult
INSERT INTO billing_metric (name, code, aggregation, field_name)
VALUES (?, ?, ?, ?);

-- name: GetBillingMetricById :one
SELECT id, name, code, aggregation, field_name
FROM billing_metric
WHERE id = ?;

-- name: UpdateBillingMetric :execresult
UPDATE billing_metric
SET name = ?, code = ?, aggregation = ?, field_name = ?
WHERE id = ?;

-- name: DeleteBillingMetric :exec
DELETE FROM billing_metric
WHERE id = ?;

-- name: ListBillingMetrics :many
SELECT id, name, code, aggregation, field_name
FROM billing_metric
ORDER BY id;