-- +goose Up

CREATE TABLE users (
id uuid Primary Key
);

-- +goose Down
DROP TABLE users;