CREATE TABLE IF NOT EXISTS games (
    id           SERIAL PRIMARY KEY,
    steam_app_id INTEGER NOT NULL UNIQUE,
    name         VARCHAR(255) NOT NULL,
    cover_url    TEXT,
    created_at   TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP NOT NULL DEFAULT NOW()
);