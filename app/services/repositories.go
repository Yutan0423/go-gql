package services

import (
	"context"
	"go-gql/graphql/model"
	"go-gql/repository"
)

func GetRepoByFullName(ctx context.Context, owner, name string) (*model.Repository, error) {
	db := repository.New(repository.Db)
	repo, err := db.GetRepoByFullName(ctx, repository.GetRepoByFullNameParams{Owner: owner, Name: name})
	if err != nil {
		return nil, err
	}

	return convertRepository(repo), nil
}

func convertRepository(repo repository.GetRepoByFullNameRow) *model.Repository {
	return &model.Repository{
		ID:        repo.ID,
		Owner:     &model.User{ID: repo.Owner},
		Name:      repo.Name,
		CreatedAt: repo.CreatedAt,
	}
}
