-- name: CreatePost :one
INSERT INTO posts (
    title, body
) VALUES (
    $1, $2
)
RETURNING id, title, created_at;
