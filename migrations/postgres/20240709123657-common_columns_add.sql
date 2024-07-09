-- +migrate Up

-- Alter football_clubs table to add isActive, createdAt, updatedAt fields
ALTER TABLE football_clubs
    ADD COLUMN is_active BOOLEAN DEFAULT TRUE,
    ADD COLUMN created_at TIMESTAMP WITH TIME ZONE,
    ADD COLUMN updated_at TIMESTAMP WITH TIME ZONE;

-- Update existing rows to set default values
UPDATE football_clubs
SET created_at = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP;

-- Alter players table to add isActive, createdAt, updatedAt fields
ALTER TABLE players
    ADD COLUMN is_active BOOLEAN DEFAULT TRUE,
    ADD COLUMN created_at TIMESTAMP WITH TIME ZONE,
    ADD COLUMN updated_at TIMESTAMP WITH TIME ZONE;

-- Update existing rows to set default values
UPDATE players
SET created_at = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP;

-- +migrate Down

-- Revert changes made in Up migration for football_clubs table
ALTER TABLE football_clubs
DROP COLUMN is_active,
    DROP COLUMN created_at,
    DROP COLUMN updated_at;

-- Revert changes made in Up migration for players table
ALTER TABLE players
DROP COLUMN is_active,
    DROP COLUMN created_at,
    DROP COLUMN updated_at;
