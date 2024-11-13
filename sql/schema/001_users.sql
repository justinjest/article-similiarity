-- +goose Up

CREATE TABLE users (
id uuid Primary Key
);
CREATE TABLE wordsInSet(
    id uuid PRIMARY KEY,
    word TEXT NOT NULL,
    inDocs INTEGER NOT NULL
);

-- +goose Down
DROP TABLE wordsInSet;
DROP TABLE users;