-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
    ADD COLUMN api_key VARCHAR(64);

CREATE TRIGGER set_default_api_key
    AFTER INSERT ON users
    BEGIN
        -- Check if the new row's api_key is NULL
        UPDATE users
        SET api_key = lower(hex(randomblob(32)))
        WHERE NEW.api_key IS NULL;
    END;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER set_default_api_key;

ALTER TABLE users
    DROP COLUMN api_key;
-- +goose StatementEnd
