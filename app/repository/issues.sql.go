// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: issues.sql

package repository

import (
	"context"
	"database/sql"
)

const addTestIssues = `-- name: AddTestIssues :execresult
INSERT INTO issues(id, url, title, closed, number, repository) VALUES
('ISSUE_1', 'http://example.com/repo1/issue/1', 'First Issue', 1, 1, 'REPO_1'),
('ISSUE_2', 'http://example.com/repo1/issue/2', 'Second Issue', 0, 2, 'REPO_1'),
('ISSUE_3', 'http://example.com/repo1/issue/3', 'Third Issue', 0, 3, 'REPO_1')
`

func (q *Queries) AddTestIssues(ctx context.Context) (sql.Result, error) {
	return q.db.ExecContext(ctx, addTestIssues)
}

const getIssueByRepoAndNumber = `-- name: GetIssueByRepoAndNumber :one
SELECT id, url, title, closed, number, repository
FROM issues
WHERE repository = ?
AND number = ?
`

type GetIssueByRepoAndNumberParams struct {
	Repository string
	Number     int32
}

func (q *Queries) GetIssueByRepoAndNumber(ctx context.Context, arg GetIssueByRepoAndNumberParams) (Issue, error) {
	row := q.db.QueryRowContext(ctx, getIssueByRepoAndNumber, arg.Repository, arg.Number)
	var i Issue
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.Title,
		&i.Closed,
		&i.Number,
		&i.Repository,
	)
	return i, err
}
