// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: projects.sql

package repository

import (
	"context"
	"database/sql"
)

const addTestProject = `-- name: AddTestProject :execresult
INSERT INTO projects(id, title, url, owner) VALUES
('PJ_1', 'My Project', 'http://example.com/project/1', 'U_1')
`

func (q *Queries) AddTestProject(ctx context.Context) (sql.Result, error) {
	return q.db.ExecContext(ctx, addTestProject)
}
