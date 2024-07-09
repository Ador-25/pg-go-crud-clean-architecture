-- +migrate Up

ALTER TABLE football_clubs
    ALTER COLUMN id SET DEFAULT nextval('football_clubs_id_seq');

ALTER TABLE players
    ALTER COLUMN id SET DEFAULT nextval('players_id_seq');

-- +migrate Down

ALTER TABLE football_clubs
    ALTER COLUMN id DROP DEFAULT;

ALTER TABLE players
    ALTER COLUMN id DROP DEFAULT;
