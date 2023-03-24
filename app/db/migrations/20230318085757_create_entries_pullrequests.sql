-- migrate:up
CREATE TABLE IF NOT EXISTS `pullrequests`(
    id VARCHAR(255) PRIMARY KEY NOT NULL,
    base_ref_name VARCHAR(255) NOT NULL,
    closed INT NOT NULL DEFAULT 0,
    head_ref_name VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL,
    number INT NOT NULL,
    repository VARCHAR(255) NOT NULL,
    CHECK (closed IN (0, 1)),
    FOREIGN KEY (repository) REFERENCES repositories(id)
);

-- migrate:down
DROP TABLE `pullrequests`;

