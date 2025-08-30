-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeedFromURL :one
SELECT * from feeds
WHERE url = $1;

-- name: GetFeeds :many
SELECT * from feeds;

