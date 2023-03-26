-- name: FetchUser :many
SELECT *
FROM users;

-- name: AddTestUser :execresult
INSERT INTO users(id, name) VALUES 
('U_1', 'hsaki');

-- name: FetchUserByName :one
SELECT
  id,
  name
FROM users
WHERE name = ?
LIMIT 1;

-- name: FetchUserByID :one
SELECT
  id,
  name
FROM users
WHERE id = ?
LIMIT 1;