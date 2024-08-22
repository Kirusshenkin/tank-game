-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS players (
    id VARCHAR(255) PRIMARY KEY,
    x INT NOT NULL,
    y INT NOT NULL,
    direction INT NOT NULL
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS players;