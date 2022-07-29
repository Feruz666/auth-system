-- name: CreateLayer :one
INSERT INTO layers (
  username,
  date,
  layer
) VALUES (
  $1, $2, $3
) RETURNING *;


-- name: ListLayer :many
SELECT * FROM layers
ORDER BY id
LIMIT $1
OFFSET $2;