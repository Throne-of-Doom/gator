-- name: CreateFeed :one
INSERT INTO feeds (name, url, user_id)
VALUES (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: ListFeeds :many
SELECT feeds.name AS feed_name, feeds.url AS feed_url, users.name AS user_name 
FROM feeds
INNER JOIN users
ON feeds.user_id = users.id;
