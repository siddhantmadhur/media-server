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

-- name: UpdateUser :exec
UPDATE profiles 
SET username = ?, password = ?
WHERE id = ?; 

-- name: GetUserFromUsername :one
SELECT * FROM profiles 
WHERE username = ?; 

-- name: CreateSession :one
INSERT INTO sessions (id, user_id, created_at, access_token, refresh_token, refresh_expires_at, access_expires_at, device, device_name, client_name, client_version)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING *; 

-- name: CreateMediaLibrary :one
INSERT INTO media_library(created_at, name, description, device_path, media_type, owner_id) 
VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: GetAllMediaLibraries :many
SELECT * FROM media_library;

-- name: AddNewContentFile :exec
INSERT INTO content_library (media_library_id,
created_at,
file_path,
extension,
name,
title,
description,
cover_url,
season_no,
episode_no) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetContentInfo :one
SELECT * FROM content_library
left join media_library
on content_library.media_library_id = media_library.id
where content_library.id = ?;

-- name: GetMediaLibrary :one
SELECT * FROM media_library
WHERE id = ?;

-- name: GetAllContentFiles :many
SELECT * FROM content_library
LEFT JOIN media_library
ON media_library.id = content_library.media_library_id
WHERE media_library_id = ?;

