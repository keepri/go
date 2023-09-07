-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name TEXT NOT NULL,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_name ON users(name);

CREATE TRIGGER set_updated_at_on_update
    AFTER UPDATE ON users
    FOR EACH ROW
    BEGIN
        UPDATE users
        SET updated_at = CURRENT_TIMESTAMP
        WHERE id = NEW.id;
    END;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER set_updated_at_on_update;

DROP INDEX idx_name;

DROP TABLE users;
-- +goose StatementEnd
