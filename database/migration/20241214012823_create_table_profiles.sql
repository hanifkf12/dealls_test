-- +goose Up
-- +goose StatementBegin
CREATE TABLE profiles (
                          id INT AUTO_INCREMENT PRIMARY KEY,
                          user_id INT NOT NULL,
                          name VARCHAR(255) NOT NULL,
                          avatar VARCHAR(255) DEFAULT NULL,
                          age INT NOT NULL,
                          gender ENUM('male', 'female') NOT NULL,
                          bio TEXT DEFAULT NULL,
                          location VARCHAR(255) DEFAULT NULL,
                          created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          deleted_at DATETIME DEFAULT NULL,
                          FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS profiles;
-- +goose StatementEnd
