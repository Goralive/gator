-- name: CreateFead :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id) 
VALUES ($1,$2,$3,$4,$5,$6)
RETURNING *;

-- name: GetFeeds :many
SELECT f.name, f.url, s.name FROM feeds as f
JOIN users as s ON f.user_id = s.id
ORDER BY s.id;
