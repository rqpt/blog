-- name: CreatePost :one
INSERT INTO posts (
    title, body
) VALUES (
    $1, $2
)
RETURNING id, title, created_at;

-- name: ShowPost :one
SELECT * FROM posts
WHERE id = $1;
