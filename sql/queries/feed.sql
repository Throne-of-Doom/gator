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

-- name: CreateFeedFollow :many
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (user_id, feed_id)
    VALUES (
        $1,
        $2
    )
    RETURNING * )
SELECT 
    inserted_feed_follow.*,
    users.name AS user_name,
    feeds.name AS feed_name
    FROM inserted_feed_follow
    INNER JOIN users
    ON inserted_feed_follow.user_id = users.id
    INNER JOIN feeds
    ON inserted_feed_follow.feed_id = feeds.id;

-- name: GetFeedByURL :one
SELECT id, name, url
FROM feeds
WHERE url = $1;

-- name: GetFeedFollowsForUser :many
SELECT feeds.name AS feed_name, users.name AS user_name
FROM feed_follows
INNER JOIN users ON feed_follows.user_id = users.id
INNER JOIN feeds ON feed_follows.feed_id = feeds.id
WHERE feed_follows.user_id = $1;

-- name: DeleteFeedFollow :one
DELETE FROM feed_follows
USING feeds
WHERE feed_follows.user_id = $1 AND feed_follows.feed_id = feeds.id AND feeds.url = $2
RETURNING feed_follows.id, feed_follows.user_id, feed_follows.feed_id;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = CURRENT_TIMESTAMP, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT last_fetched_at, id, url
FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;