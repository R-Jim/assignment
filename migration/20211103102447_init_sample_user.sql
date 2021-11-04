-- +goose Up
-- +goose StatementBegin
SELECT
    'up SQL query';

-- +goose StatementEnd
INSERT INTO
    `twitter`.`users` (
        `created_at`,
        `updated_at`,
        `username`,
        `password`
    )
VALUES
    (NOW(), NOW(), 'user', '123');

-- +goose Down
-- +goose StatementBegin
SELECT
    'down SQL query';

-- +goose StatementEnd