// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package repository

import (
	"database/sql"
	"time"
)

type Issue struct {
	ID         string
	Url        string
	Title      string
	Closed     int32
	Number     int32
	Repository string
}

type Project struct {
	ID    string
	Title string
	Url   string
	Owner string
}

type Projectcard struct {
	ID          string
	Project     string
	Issue       sql.NullString
	Pullrequest sql.NullString
}

type Pullrequest struct {
	ID          string
	BaseRefName string
	Closed      int32
	HeadRefName string
	Url         string
	Number      int32
	Repository  string
}

type Repository struct {
	ID        string
	Owner     string
	Name      string
	CreatedAt time.Time
}

type User struct {
	ID        string
	Name      string
	ProjectV2 sql.NullString
}
