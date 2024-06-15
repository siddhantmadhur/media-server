CREATE TABLE IF NOT EXISTS profiles (
    id INTEGER PRIMARY KEY,
    username text NOT NULL,
    password BLOB,
    type INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS sessions (
    id TEXT PRIMARY KEY,
    user_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    device TEXT NOT NULL,
    device_name TEXT NOT NULL,
    client_name TEXT NOT NULL,
    client_version TEXT NOT NULL,

    FOREIGN KEY(user_id) REFERENCES profiles(id)
);

CREATE TABLE IF NOT EXISTS media_library (
    id INTEGER PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    device_path TEXT NOT NULL,
    media_type TEXT NOT NULL,
    owner_id INTEGER NOT NULL,

    FOREIGN KEY(owner_id) REFERENCES profiles(id),
    UNIQUE(device_path)
);

CREATE TABLE IF NOT EXISTS content_library (
    id INTEGER PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    file_path TEXT NOT NULL,
    media_library_id INTEGER NOT NULL,
    extension TEXT NOT NULL,
    name TEXT NOT NULL,

    FOREIGN KEY(media_library_id) REFERENCES media_library(id),
    UNIQUE(file_path)
);

CREATE TABLE IF NOT EXISTS content_metadata (
    id INTEGER PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    content_id INTEGER NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    poster_url TEXT NOT NULL,
    release_date TIMESTAMP NOT NULL,
    season_number INTEGER,
    episode_number INTEGER,
    type TEXT NOT NULL,
    
    FOREIGN KEY(content_id) REFERENCES content_library(id),
    UNIQUE(content_id)
);

