-- name: GetAuthor :one
SELECT * FROM blog_user
WHERE email = $1 LIMIT 1;

-- name: ListAuthor :many
SELECT * FROM blog_user
ORDER BY id;

-- name: CreateAuthor :one
INSERT INTO blog_user (
    email, username, name, password
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM blog_user
WHERE id = $1;