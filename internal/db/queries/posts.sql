-- name: PostCreate :one
INSERT INTO posts (title, body)
VALUES ($1, $2)
RETURNING id, title, slug, created_at;

-- name: PostIndex :many
SELECT * FROM posts;

-- name: PostShow :one
SELECT * FROM posts
WHERE id = $1;
