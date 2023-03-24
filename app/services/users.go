package services

import (
	"context"
	"go-gql/graphql/model"
	"go-gql/repository"
	"log"
)

func convertUser(user repository.FetchUserByNameRow) *model.User {
	return &model.User{
		ID:   user.ID,
		Name: user.Name,
	}
}

func GetUserByName(ctx context.Context, name string) (*model.User, error) {
	Db := repository.New(repository.Db)
	user, err := Db.FetchUserByName(ctx, name)
	if err != nil {
		log.Fatalf("failed to fetch user by name: %v", err)
	}

	return convertUser(user), nil
}
