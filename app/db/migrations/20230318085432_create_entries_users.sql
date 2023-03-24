-- migrate:up
CREATE TABLE IF NOT EXISTS `users`(
    id VARCHAR(255) PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    project_v2 TEXT
);

-- migrate:down
DROP TABLE `users`;
