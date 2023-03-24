package services

import (
	"context"
	"fmt"
	"go-gql/repository"
)

func AddDummyData() {
	Db := repository.New(repository.Db)
	context := context.Background()
	user, err := Db.AddTestUser(context)
	fmt.Println(user, err)
	project, err := Db.AddTestProject(context)
	fmt.Println(project, err)
	issues, err := Db.AddTestIssues(context)
	fmt.Println(issues, err)
	repositories, err := Db.AddTestRepository(context)
	fmt.Println(repositories, err)
}
