CREATE TABLE IF NOT EXISTS profiles (
    id INTEGER PRIMARY KEY,
    username text NOT NULL,
    password text NOT NULL,
    type INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS sessions (
    id INTEGER PRIMARY KEY,
    user_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    refresh_token text NOT NULL,

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
