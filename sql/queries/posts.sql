-- name: CreatePost :one
INSERT INTO posts (id,created_at,updated_at,title,url,description,published_at,feed_id)
VALUES(
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
RETURNING *;

-- name: GetPostsForUser :many
WITH feed_follows as (
    SELECT feed_id FROM feed_follows f WHERE f.user_id = $1
)
SELECT * from posts p WHERE p.feed_id in (
    SELECT feed_id FROM feed_follows
) ORDER BY p.published_at DESC NULLS LAST LIMIT $2;
