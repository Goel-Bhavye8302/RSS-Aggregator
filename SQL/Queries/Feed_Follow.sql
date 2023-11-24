-- name: CreateFeedFollow :one
INSERT INTO Feed_follows (id, created_at, updated_at, user_id, feed_id) 
VALUES ($1, $2, $3, $4, $5)
RETURNING *;


-- name: GetFeedFollows :many 
SELECT * FROM Feed_follows WHERE user_id = $1;

-- name: DeleteFeedFollow :exec 
DELETE FROM Feed_follows WHERE feed_id = $1 AND user_id = $2;