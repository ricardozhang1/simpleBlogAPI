-- name: GetTag :one
SELECT * FROM blog_tag
WHERE id = $1 LIMIT 1;

-- name: ListTag :many
SELECT * FROM blog_tag
where id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: CreateTag :one
INSERT INTO blog_tag (
    name, created_id
) VALUES (
    $1, $2
)
RETURNING *;

-- name: UpdateTag :one
UPDATE blog_tag SET name=$1, modified_id=$2, modified_at=$3
WHERE id=$4
RETURNING *;;

-- name: DeleteTag :exec
DELETE FROM blog_tag
WHERE id = $1;