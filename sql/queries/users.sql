-- name: CreateUser :one
INSERT INTO users (id, created_ad, updated_at, name)
VALUES ($1, $2, $3, $4)
RETURNING *;