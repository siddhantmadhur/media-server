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

-- name: CreateMediaLibrary :one
INSERT INTO media_library(created_at, name, description, device_path, media_type, owner_id) 
VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: AddContent :one
INSERT INTO content_library(created_at, file_path, media_library_id, extension, name)
VALUES ( ?, ?, ?, ? , ?)
RETURNING *;

-- name: LinkContentMetadata :one
INSERT INTO content_metadata(created_at, content_id, title, description, poster_url, release_date)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: GetMediaLibrary :one
SELECT * FROM media_library
WHERE id = ?;
