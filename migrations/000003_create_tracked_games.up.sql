CREATE TABLE IF NOT EXISTS tracked_games (
    id           SERIAL PRIMARY KEY,
    user_id      INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    game_id      INTEGER NOT NULL REFERENCES games(id) ON DELETE CASCADE,
    target_price NUMERIC(10,2) NULL,
    created_at   TIMESTAMP NOT NULL DEFAULT NOW(),

    UNIQUE (user_id, game_id)
);