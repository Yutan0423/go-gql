-- migrate:up
CREATE TABLE IF NOT EXISTS `issues`(
    id VARCHAR(255) PRIMARY KEY NOT NULL,
    url VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    closed INT NOT NULL DEFAULT 0,
    number INT NOT NULL,
    repository VARCHAR(255) NOT NULL,
    CHECK (closed IN (0, 1)),
    FOREIGN KEY (repository) REFERENCES repositories(id)
);

-- migrate:down
DROP TABLE `issues`;

