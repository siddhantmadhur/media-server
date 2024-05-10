
CREATE TABLE IF NOT EXISTS profiles (
    id INTEGER PRIMARY KEY,
    username text NOT NULL,
    password text NOT NULL,
    type INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS settings (
    key text PRIMARY KEY,
    value text NOT NULL
);
