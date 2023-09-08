-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
    ADD COLUMN api_key VARCHAR(64) NOT NULL DEFAULT (lower(hex(randomblob(32))));

CREATE UNIQUE INDEX unique_api_key ON users(api_key);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX unique_api_key;

ALTER TABLE users
    DROP COLUMN api_key;
-- +goose StatementEnd
