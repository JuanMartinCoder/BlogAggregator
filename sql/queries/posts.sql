-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at,title, url, description, published_at, feed_id)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING *;

-- name: GetPostsForUser :many
SELECT * FROM posts INNER JOIN feeds ON feeds.id = posts.feed_id 
WHERE posts.feed_id IN (SELECT feed_id FROM feed_follows f WHERE f.user_id = $1) 
ORDER BY published_at DESC LIMIT $2;