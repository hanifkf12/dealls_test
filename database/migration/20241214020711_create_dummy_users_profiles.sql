-- +goose Up
-- +goose StatementBegin
-- Insert 12 dummy users
INSERT INTO users (email, password, created_at, updated_at) VALUES
                                                                ('john.doe@example.com', '$2a$10$E5nTCRdbU2O9RlSfXe9MOeQtsPIZcfEp69USFKs8DdD8ZRhg6uKFO', NOW(), NOW()),
                                                                ('jane.doe@example.com', '$2a$10$E5nTCRdbU2O9RlSfXe9MOeQtsPIZcfEp69USFKs8DdD8ZRhg6uKFO', NOW(), NOW()),
                                                                ('alice.smith@example.com', '$2a$10$E5nTCRdbU2O9RlSfXe9MOeQtsPIZcfEp69USFKs8DdD8ZRhg6uKFO', NOW(), NOW()),
                                                                ('bob.johnson@example.com', '$2a$10$E5nTCRdbU2O9RlSfXe9MOeQtsPIZcfEp69USFKs8DdD8ZRhg6uKFO', NOW(), NOW()),
                                                                ('charlie.brown@example.com', '$2a$10$E5nTCRdbU2O9RlSfXe9MOeQtsPIZcfEp69USFKs8DdD8ZRhg6uKFO', NOW(), NOW()),
                                                                ('daisy.jones@example.com', '$2a$10$E5nTCRdbU2O9RlSfXe9MOeQtsPIZcfEp69USFKs8DdD8ZRhg6uKFO', NOW(), NOW()),
                                                                ('emma.watson@example.com', '$2a$10$E5nTCRdbU2O9RlSfXe9MOeQtsPIZcfEp69USFKs8DdD8ZRhg6uKFO', NOW(), NOW()),
                                                                ('frank.castle@example.com', '$2a$10$E5nTCRdbU2O9RlSfXe9MOeQtsPIZcfEp69USFKs8DdD8ZRhg6uKFO', NOW(), NOW()),
                                                                ('grace.hopper@example.com', '$2a$10$E5nTCRdbU2O9RlSfXe9MOeQtsPIZcfEp69USFKs8DdD8ZRhg6uKFO', NOW(), NOW()),
                                                                ('henry.ford@example.com', '$2a$10$E5nTCRdbU2O9RlSfXe9MOeQtsPIZcfEp69USFKs8DdD8ZRhg6uKFO', NOW(), NOW()),
                                                                ('isabella.rossi@example.com', '$2a$10$E5nTCRdbU2O9RlSfXe9MOeQtsPIZcfEp69USFKs8DdD8ZRhg6uKFO', NOW(), NOW()),
                                                                ('jack.daniels@example.com', '$2a$10$E5nTCRdbU2O9RlSfXe9MOeQtsPIZcfEp69USFKs8DdD8ZRhg6uKFO', NOW(), NOW());

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
