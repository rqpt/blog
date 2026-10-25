-- name: IndexPosts :many
SELECT * FROM posts;

-- name: ShowPost :one
SELECT * FROM posts
WHERE id = $1;

-- name: CreatePost :one
INSERT INTO posts (title, body, frontmatter)
VALUES ($1, $2, $3)
RETURNING id, title, slug, created_at;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1;
