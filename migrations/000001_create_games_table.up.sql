CREATE TABLE IF NOT EXISTS games (
    id          SERIAL PRIMARY KEY,
    name        TEXT NOT NULL,
    description TEXT,
    genre       TEXT NOT NULL,
    platform    TEXT NOT NULL,
    price       DECIMAL(10,2),
    rating      DECIMAL(3,1),
    cover_url   TEXT,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW()
);