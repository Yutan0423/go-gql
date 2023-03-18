-- name: FetchUser :many
SELECT *
FROM users;

-- name: RegisterUser :execresult
INSERT INTO users (uid, name, email, image) VALUES (?, ?, ?, ?);

-- name: FetchCourses :many
SELECT *
FROM clear_courses;

-- name: FetchProblem :many
SELECT *
FROM clear_problems;