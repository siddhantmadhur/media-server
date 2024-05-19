-- name: GetProfiles :many
SELECT id, username FROM profiles;

-- name: IsFinishedSetup :one 
SELECT count(*) FROM profiles;

-- name: CreateProfile :exec
INSERT INTO profiles (username, password, type) 
VALUES ( ?, ?, ? );

-- name: GetAdminUser :one
SELECT * FROM profiles
WHERE type = 0;

-- name: UpdateAdminUser :exec
UPDATE profiles 
SET username = ?, password = ?
WHERE type = 0;

-- name: GetUserWithPassword :one
SELECT * FROM profiles 
WHERE username = ? and password = ?; 

-- name: CreateSession :one
INSERT INTO sessions (id, user_id, created_at, expires_at, device, device_name, client_name, client_version)
VALUES (?, ?, ?, ?, ?, ?, ?, ?)
RETURNING *;
