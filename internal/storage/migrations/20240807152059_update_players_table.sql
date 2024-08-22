-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE players
ADD COLUMN created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
ADD COLUMN name VARCHAR(255),
ADD COLUMN device_id VARCHAR(255);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE players
DROP COLUMN created_at,
DROP COLUMN name,
DROP COLUMN device_id;