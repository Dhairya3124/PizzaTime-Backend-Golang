--+goose Up
ALTER TABLE player ADD COLUMN date_updated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;
--+goose Down
ALTER TABLE player DROP COLUMN date_updated;
