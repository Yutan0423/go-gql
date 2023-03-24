-- migrate:up
CREATE TABLE IF NOT EXISTS `repositories`(
    id VARCHAR(255) PRIMARY KEY NOT NULL,
    owner VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (owner) REFERENCES users(id)
);

-- migrate:down
DROP TABLE `repositories`;
