CREATE TABLE IF NOT EXISTS price_history (
    id               SERIAL PRIMARY KEY,
    game_id          INTEGER NOT NULL REFERENCES games(id) ON DELETE CASCADE,
    lang             VARCHAR(2) NOT NULL,
    price            NUMERIC(10, 2) NOT NULL,
    currency         VARCHAR(3) NOT NULL,
    discount_percent INTEGER NOT NULL DEFAULT 0,
    checked_at       TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_price_history_game_id_cc ON price_history (game_id, lang);
CREATE INDEX idx_price_history_checked_at ON price_history (checked_at);