-- name: CreateEvent :execresult
INSERT INTO event (customer_id, code, timestamp, value)
VALUES (?, ?, ?, ?);

-- name: GetEventById :many
SELECT id, customer_id, code, timestamp, value
FROM event
WHERE id = ?;

-- name: SumEventsByCustomerIdAndCodeName :one
SELECT customer_id, code, CAST(SUM(value) AS FLOAT) AS sum
FROM event
WHERE customer_id = ? AND code = ?
GROUP BY customer_id, code;

-- name: CountEventsByCustomerIdAndCode :one
SELECT customer_id, code, COUNT(*) AS count
FROM event
WHERE customer_id = ? AND code = ?
GROUP BY customer_id, code;

-- name: SumEventsByCustomerIdCodeNameStartRangeAndEndRange :one
SELECT customer_id, code, CAST(SUM(value) AS FLOAT) AS sum
FROM event
WHERE customer_id = ? AND code = ? AND value <= ? AND value >= ?
GROUP BY customer_id, code;
