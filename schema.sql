CREATE TABLE IF NOT EXISTS profiles (
    id INTEGER PRIMARY KEY,
    username text NOT NULL,
    password text NOT NULL,
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

CREATE TABLE IF NOT EXISTS media_libraries (
    id INTEGER PRIMARY KEY,
    owner INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL,
    path TEXT NOT NULL,
    type TEXT NOT NULL,

    FOREIGN KEY(owner) REFERENCES profiles(id)
)
