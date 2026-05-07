CREATE TABLE IF NOT EXISTS users (
    id          SERIAL PRIMARY KEY,
    api_key     VARCHAR(64) NOT NULL UNIQUE,
    lang        VARCHAR(2) NOT NULL DEFAULT 'US',
    email       VARCHAR(255) UNIQUE,
    telegram_id BIGINT UNIQUE ,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW()
);