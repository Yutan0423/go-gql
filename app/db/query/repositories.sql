-- name: AddTestRepository :execresult
INSERT INTO repositories(id, owner, name) VALUES
('REPO_1', 'U_1', 'repo1');

-- name: GetRepoByFullName :one
SELECT id, name, owner, created_at
FROM repositories
WHERE name = ?
AND owner = ?