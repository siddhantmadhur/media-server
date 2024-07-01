// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package storage

import (
	"context"
	"database/sql"
	"time"
)

const addNewContentFile = `-- name: AddNewContentFile :one
INSERT INTO content_library (
    media_library_id,
    created_at,
    file_path,
    name,
    media_title,
    description,
    cover_url,
    parent_id,
    external_provider,
    external_provider_id,
    media_type,
    classifier
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING id, media_library_id, created_at, file_path, name, media_title, description, cover_url, parent_id, classifier, media_type, external_provider, external_provider_id
`

type AddNewContentFileParams struct {
	MediaLibraryID     int64
	CreatedAt          time.Time
	FilePath           string
	Name               string
	MediaTitle         string
	Description        sql.NullString
	CoverUrl           sql.NullString
	ParentID           sql.NullInt64
	ExternalProvider   sql.NullString
	ExternalProviderID sql.NullInt64
	MediaType          string
	Classifier         string
}

func (q *Queries) AddNewContentFile(ctx context.Context, arg AddNewContentFileParams) (ContentLibrary, error) {
	row := q.db.QueryRowContext(ctx, addNewContentFile,
		arg.MediaLibraryID,
		arg.CreatedAt,
		arg.FilePath,
		arg.Name,
		arg.MediaTitle,
		arg.Description,
		arg.CoverUrl,
		arg.ParentID,
		arg.ExternalProvider,
		arg.ExternalProviderID,
		arg.MediaType,
		arg.Classifier,
	)
	var i ContentLibrary
	err := row.Scan(
		&i.ID,
		&i.MediaLibraryID,
		&i.CreatedAt,
		&i.FilePath,
		&i.Name,
		&i.MediaTitle,
		&i.Description,
		&i.CoverUrl,
		&i.ParentID,
		&i.Classifier,
		&i.MediaType,
		&i.ExternalProvider,
		&i.ExternalProviderID,
	)
	return i, err
}

const createMediaLibrary = `-- name: CreateMediaLibrary :one
INSERT INTO media_library(created_at, name, description, device_path, media_type, owner_id) 
VALUES (?, ?, ?, ?, ?, ?)
RETURNING id, created_at, name, description, device_path, media_type, owner_id
`

type CreateMediaLibraryParams struct {
	CreatedAt   time.Time
	Name        string
	Description string
	DevicePath  string
	MediaType   string
	OwnerID     int64
}

func (q *Queries) CreateMediaLibrary(ctx context.Context, arg CreateMediaLibraryParams) (MediaLibrary, error) {
	row := q.db.QueryRowContext(ctx, createMediaLibrary,
		arg.CreatedAt,
		arg.Name,
		arg.Description,
		arg.DevicePath,
		arg.MediaType,
		arg.OwnerID,
	)
	var i MediaLibrary
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Name,
		&i.Description,
		&i.DevicePath,
		&i.MediaType,
		&i.OwnerID,
	)
	return i, err
}

const createProfile = `-- name: CreateProfile :exec
INSERT INTO profiles (username, password, type) 
VALUES ( ?, ?, ? )
`

type CreateProfileParams struct {
	Username string
	Password []byte
	Type     int64
}

func (q *Queries) CreateProfile(ctx context.Context, arg CreateProfileParams) error {
	_, err := q.db.ExecContext(ctx, createProfile, arg.Username, arg.Password, arg.Type)
	return err
}

const createSession = `-- name: CreateSession :one
INSERT INTO sessions (id, user_id, created_at, access_token, refresh_token, refresh_expires_at, access_expires_at, device, device_name, client_name, client_version)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING id, user_id, access_token, refresh_token, created_at, access_expires_at, refresh_expires_at, device, device_name, client_name, client_version
`

type CreateSessionParams struct {
	ID               string
	UserID           int64
	CreatedAt        time.Time
	AccessToken      string
	RefreshToken     string
	RefreshExpiresAt time.Time
	AccessExpiresAt  time.Time
	Device           string
	DeviceName       string
	ClientName       string
	ClientVersion    string
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error) {
	row := q.db.QueryRowContext(ctx, createSession,
		arg.ID,
		arg.UserID,
		arg.CreatedAt,
		arg.AccessToken,
		arg.RefreshToken,
		arg.RefreshExpiresAt,
		arg.AccessExpiresAt,
		arg.Device,
		arg.DeviceName,
		arg.ClientName,
		arg.ClientVersion,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.AccessToken,
		&i.RefreshToken,
		&i.CreatedAt,
		&i.AccessExpiresAt,
		&i.RefreshExpiresAt,
		&i.Device,
		&i.DeviceName,
		&i.ClientName,
		&i.ClientVersion,
	)
	return i, err
}

const getAdminUser = `-- name: GetAdminUser :one
SELECT id, username, password, type FROM profiles
WHERE type = 0
`

func (q *Queries) GetAdminUser(ctx context.Context) (Profile, error) {
	row := q.db.QueryRowContext(ctx, getAdminUser)
	var i Profile
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Type,
	)
	return i, err
}

const getAllContentFiles = `-- name: GetAllContentFiles :many
SELECT content_library.id, media_library_id, content_library.created_at, file_path, content_library.name, media_title, content_library.description, cover_url, parent_id, classifier, content_library.media_type, external_provider, external_provider_id, media_library.id, media_library.created_at, media_library.name, media_library.description, device_path, media_library.media_type, owner_id FROM content_library
LEFT JOIN media_library
ON media_library.id = content_library.media_library_id
WHERE media_library_id = ?
`

type GetAllContentFilesRow struct {
	ID                 int64
	MediaLibraryID     int64
	CreatedAt          time.Time
	FilePath           string
	Name               string
	MediaTitle         string
	Description        sql.NullString
	CoverUrl           sql.NullString
	ParentID           sql.NullInt64
	Classifier         string
	MediaType          string
	ExternalProvider   sql.NullString
	ExternalProviderID sql.NullInt64
	ID_2               sql.NullInt64
	CreatedAt_2        sql.NullTime
	Name_2             sql.NullString
	Description_2      sql.NullString
	DevicePath         sql.NullString
	MediaType_2        sql.NullString
	OwnerID            sql.NullInt64
}

func (q *Queries) GetAllContentFiles(ctx context.Context, mediaLibraryID int64) ([]GetAllContentFilesRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllContentFiles, mediaLibraryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllContentFilesRow
	for rows.Next() {
		var i GetAllContentFilesRow
		if err := rows.Scan(
			&i.ID,
			&i.MediaLibraryID,
			&i.CreatedAt,
			&i.FilePath,
			&i.Name,
			&i.MediaTitle,
			&i.Description,
			&i.CoverUrl,
			&i.ParentID,
			&i.Classifier,
			&i.MediaType,
			&i.ExternalProvider,
			&i.ExternalProviderID,
			&i.ID_2,
			&i.CreatedAt_2,
			&i.Name_2,
			&i.Description_2,
			&i.DevicePath,
			&i.MediaType_2,
			&i.OwnerID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllMediaLibraries = `-- name: GetAllMediaLibraries :many
SELECT id, created_at, name, description, device_path, media_type, owner_id FROM media_library
`

func (q *Queries) GetAllMediaLibraries(ctx context.Context) ([]MediaLibrary, error) {
	rows, err := q.db.QueryContext(ctx, getAllMediaLibraries)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MediaLibrary
	for rows.Next() {
		var i MediaLibrary
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.Name,
			&i.Description,
			&i.DevicePath,
			&i.MediaType,
			&i.OwnerID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getContentFromExternalId = `-- name: GetContentFromExternalId :one
SELECT id, media_library_id, created_at, file_path, name, media_title, description, cover_url, parent_id, classifier, media_type, external_provider, external_provider_id FROM content_library
WHERE external_provider_id = ?
`

func (q *Queries) GetContentFromExternalId(ctx context.Context, externalProviderID sql.NullInt64) (ContentLibrary, error) {
	row := q.db.QueryRowContext(ctx, getContentFromExternalId, externalProviderID)
	var i ContentLibrary
	err := row.Scan(
		&i.ID,
		&i.MediaLibraryID,
		&i.CreatedAt,
		&i.FilePath,
		&i.Name,
		&i.MediaTitle,
		&i.Description,
		&i.CoverUrl,
		&i.ParentID,
		&i.Classifier,
		&i.MediaType,
		&i.ExternalProvider,
		&i.ExternalProviderID,
	)
	return i, err
}

const getContentFromPath = `-- name: GetContentFromPath :one
SELECT id, media_library_id, created_at, file_path, name, media_title, description, cover_url, parent_id, classifier, media_type, external_provider, external_provider_id FROM content_library
WHERE file_path = ?
`

func (q *Queries) GetContentFromPath(ctx context.Context, filePath string) (ContentLibrary, error) {
	row := q.db.QueryRowContext(ctx, getContentFromPath, filePath)
	var i ContentLibrary
	err := row.Scan(
		&i.ID,
		&i.MediaLibraryID,
		&i.CreatedAt,
		&i.FilePath,
		&i.Name,
		&i.MediaTitle,
		&i.Description,
		&i.CoverUrl,
		&i.ParentID,
		&i.Classifier,
		&i.MediaType,
		&i.ExternalProvider,
		&i.ExternalProviderID,
	)
	return i, err
}

const getContentInfo = `-- name: GetContentInfo :one
SELECT content_library.id, media_library_id, content_library.created_at, file_path, content_library.name, media_title, content_library.description, cover_url, parent_id, classifier, content_library.media_type, external_provider, external_provider_id, media_library.id, media_library.created_at, media_library.name, media_library.description, device_path, media_library.media_type, owner_id FROM content_library
left join media_library
on content_library.media_library_id = media_library.id
where content_library.id = ?
`

type GetContentInfoRow struct {
	ID                 int64
	MediaLibraryID     int64
	CreatedAt          time.Time
	FilePath           string
	Name               string
	MediaTitle         string
	Description        sql.NullString
	CoverUrl           sql.NullString
	ParentID           sql.NullInt64
	Classifier         string
	MediaType          string
	ExternalProvider   sql.NullString
	ExternalProviderID sql.NullInt64
	ID_2               sql.NullInt64
	CreatedAt_2        sql.NullTime
	Name_2             sql.NullString
	Description_2      sql.NullString
	DevicePath         sql.NullString
	MediaType_2        sql.NullString
	OwnerID            sql.NullInt64
}

func (q *Queries) GetContentInfo(ctx context.Context, id int64) (GetContentInfoRow, error) {
	row := q.db.QueryRowContext(ctx, getContentInfo, id)
	var i GetContentInfoRow
	err := row.Scan(
		&i.ID,
		&i.MediaLibraryID,
		&i.CreatedAt,
		&i.FilePath,
		&i.Name,
		&i.MediaTitle,
		&i.Description,
		&i.CoverUrl,
		&i.ParentID,
		&i.Classifier,
		&i.MediaType,
		&i.ExternalProvider,
		&i.ExternalProviderID,
		&i.ID_2,
		&i.CreatedAt_2,
		&i.Name_2,
		&i.Description_2,
		&i.DevicePath,
		&i.MediaType_2,
		&i.OwnerID,
	)
	return i, err
}

const getMediaLibrary = `-- name: GetMediaLibrary :one
SELECT id, created_at, name, description, device_path, media_type, owner_id FROM media_library
WHERE id = ?
`

func (q *Queries) GetMediaLibrary(ctx context.Context, id int64) (MediaLibrary, error) {
	row := q.db.QueryRowContext(ctx, getMediaLibrary, id)
	var i MediaLibrary
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Name,
		&i.Description,
		&i.DevicePath,
		&i.MediaType,
		&i.OwnerID,
	)
	return i, err
}

const getProfiles = `-- name: GetProfiles :many
SELECT id, username FROM profiles
`

type GetProfilesRow struct {
	ID       int64
	Username string
}

func (q *Queries) GetProfiles(ctx context.Context) ([]GetProfilesRow, error) {
	rows, err := q.db.QueryContext(ctx, getProfiles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetProfilesRow
	for rows.Next() {
		var i GetProfilesRow
		if err := rows.Scan(&i.ID, &i.Username); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserFromUsername = `-- name: GetUserFromUsername :one
SELECT id, username, password, type FROM profiles 
WHERE username = ?
`

func (q *Queries) GetUserFromUsername(ctx context.Context, username string) (Profile, error) {
	row := q.db.QueryRowContext(ctx, getUserFromUsername, username)
	var i Profile
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Type,
	)
	return i, err
}

const isFinishedSetup = `-- name: IsFinishedSetup :one
SELECT count(*) FROM profiles
`

func (q *Queries) IsFinishedSetup(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, isFinishedSetup)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE profiles 
SET username = ?, password = ?
WHERE id = ?
`

type UpdateUserParams struct {
	Username string
	Password []byte
	ID       int64
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser, arg.Username, arg.Password, arg.ID)
	return err
}
