-- name: GetArticle :one
SELECT * FROM blog_article
WHERE id = $1 LIMIT 1;

-- name: ListArticle :many
SELECT * FROM blog_article
ORDER BY id;

-- name: CreateArticle :one
INSERT INTO blog_article (
    tag_id, title, intro, content, created_id
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: DeleteArticle :exec
DELETE FROM blog_article
WHERE id = $1;