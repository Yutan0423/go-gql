// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: repositories.sql

package repository

import (
	"context"
	"database/sql"
)

const addTestRepository = `-- name: AddTestRepository :execresult
INSERT INTO repositories(id, owner, name) VALUES
('REPO_1', 'U_1', 'repo1')
`

func (q *Queries) AddTestRepository(ctx context.Context) (sql.Result, error) {
	return q.db.ExecContext(ctx, addTestRepository)
}
