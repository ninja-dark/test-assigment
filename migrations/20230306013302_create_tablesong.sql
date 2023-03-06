-- +goose Up
-- +goose StatementBegin
CREATE TABLE playlist (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    duration int NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS playlist;
-- +goose StatementEnd
