
-- +migrate Up
CREATE TABLE IF NOT EXISTS football_clubs (
    id SERIAL PRIMARY KEY,
    club_name VARCHAR(255) NOT NULL,
    sponsors TEXT[] DEFAULT ARRAY[]::TEXT[]);

CREATE TABLE IF NOT EXISTS players (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    club_id INT REFERENCES football_clubs(id) ON DELETE CASCADE);

-- +migrate Down
DROP TABLE IF EXISTS players;
DROP TABLE IF EXISTS football_clubs;