-- migrate:up
CREATE TABLE IF NOT EXISTS `projectcards`(
    id VARCHAR(255) PRIMARY KEY NOT NULL,
    project VARCHAR(255) NOT NULL,
    issue VARCHAR(255),
    pullrequest VARCHAR(255),
    FOREIGN KEY (project) REFERENCES projects(id),
    FOREIGN KEY (issue) REFERENCES issues(id),
    FOREIGN KEY (pullrequest) REFERENCES pullrequests(id),
    CHECK (issue IS NOT NULL OR pullrequest IS NOT NULL)
);

-- migrate:down
DROP TABLE `projectcards`;

