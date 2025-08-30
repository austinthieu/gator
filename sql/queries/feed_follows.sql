-- name: CreateFeedFollows :one
WITH inserted_feed_follow AS (
  INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
  VALUES ($1, $2, $3, $4, $5)
  RETURNING *
)
SELECT 
  inserted_feed_follow.*,
  feeds.name AS feed_name,
  users.name AS user_name
from inserted_feed_follow
INNER JOIN feeds ON inserted_feed_follow.feed_id = feeds.id
INNER JOIN users ON inserted_feed_follow.user_id = users.id;


-- name: GetFeedFollowsForUser :many 
SELECT
  feed_follows.id,
  feed_follows.user_id,
  feed_follows.feed_id,
  users.name AS user_name,
  feeds.name AS feed_name
FROM feed_follows
INNER JOIN users ON feed_follows.user_id = users.id
INNER JOIN feeds on feed_follows.feed_id = feeds.id
WHERE feed_follows.user_id = $1;

