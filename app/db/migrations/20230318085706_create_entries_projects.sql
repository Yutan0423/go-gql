-- migrate:up
CREATE TABLE IF NOT EXISTS `projects`(
    id VARCHAR(255) PRIMARY KEY NOT NULL,
    title VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL,
    owner VARCHAR(255) NOT NULL,
    FOREIGN KEY (owner) REFERENCES users(id)
);

-- migrate:down
DROP TABLE `projects`;

