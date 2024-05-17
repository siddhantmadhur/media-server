-- name: GetProfiles :many
SELECT id, username FROM profiles;

-- name: IsFinishedSetup :one 
SELECT count(*) FROM profiles;

-- name: CreateProfile :exec
INSERT INTO profiles (id, username, password, type) 
VALUES ( ?, ?, ?, ? );

-- name: GetUserWithPassword :one
SELECT * FROM profiles 
WHERE username = ? and password = ?; 

-- name: CreateSession :one
INSERT INTO sessions (user_id, created_at, expires_at, refresh_token)
VALUES (?, ?, ?, ?)
RETURNING *;
