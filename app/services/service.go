package services

import (
	"context"
	"go-gql/graphql/model"
	"go-gql/repository"
	"log"
)

func GetIssueByRepoAndNumber(ctx context.Context, repoID string, number int) (*model.Issue, error) {
	db := repository.New(repository.Db)
	issue, err := db.GetIssueByRepoAndNumber(ctx, repository.GetIssueByRepoAndNumberParams{Repository: repoID, Number: int32(number)})
	if err != nil {
		log.Fatalf("failed to fetch issue by repoID and number: %v", err)
	}

	return convertIssue(issue), nil
}

func convertIssue(issue repository.Issue) *model.Issue {
	// var issueURL model.MyURL
	// 参照渡しでデータ構造を変える
	// issueURL.UnmarshalGQL(issue.Url)
	// fmt.Println(issue.Url)
	// fmt.Println(issueURL)

	return &model.Issue{
		ID:         issue.ID,
		Title:      issue.Title,
		Closed:     (issue.Closed == 1),
		Number:     int(issue.Number),
		Repository: &model.Repository{ID: issue.Repository},
	}
}
