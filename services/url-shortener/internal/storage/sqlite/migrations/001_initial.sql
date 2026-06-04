CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE COLLATE NOCASE,
    email TEXT NOT NULL UNIQUE COLLATE NOCASE,
    password_hash TEXT NOT NULL,
    created_at TEXT NOT NULL,
    modified_at TEXT NOT NULL
);

CREATE TABLE sessions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token_hash TEXT NOT NULL UNIQUE,
    created_at TEXT NOT NULL,
    expires_at TEXT NOT NULL,
    last_seen_at TEXT
);

CREATE INDEX idx_sessions_user_id ON sessions(user_id);
CREATE INDEX idx_sessions_expires_at ON sessions(expires_at);

CREATE TABLE shortened_urls (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    original_url TEXT NOT NULL,
    normalized_url TEXT NOT NULL,
    created_at TEXT NOT NULL,
    modified_at TEXT NOT NULL,
    UNIQUE(user_id, normalized_url)
);

CREATE INDEX idx_shortened_urls_user_id ON shortened_urls(user_id);
CREATE INDEX idx_shortened_urls_normalized_url ON shortened_urls(normalized_url);

CREATE TABLE short_codes (
    code TEXT PRIMARY KEY,
    shortened_url_id INTEGER NOT NULL REFERENCES shortened_urls(id) ON DELETE CASCADE,
    kind TEXT NOT NULL CHECK (kind IN ('system', 'custom')),
    created_at TEXT NOT NULL,
    UNIQUE(shortened_url_id, kind)
);

CREATE INDEX idx_short_codes_shortened_url_id ON short_codes(shortened_url_id);
