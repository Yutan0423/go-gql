package services

import (
	"context"
	"go-gql/graphql/model"
	"go-gql/repository"
	"log"
)

func convertUser(user any) *model.User {
	switch user := user.(type) {
	case repository.FetchUserByNameRow:
		return &model.User{
			ID:   user.ID,
			Name: user.Name,
		}
	case repository.FetchUserByIDRow:
		return &model.User{
			ID:   user.ID,
			Name: user.Name,
		}
	default:
		return &model.User{}
	}
}

func GetUserByName(ctx context.Context, name string) (*model.User, error) {
	db := repository.New(repository.Db)
	user, err := db.FetchUserByName(ctx, name)
	if err != nil {
		log.Fatalf("failed to fetch user by name: %v", err)
	}

	return convertUser(user), nil
}

func GetUserByID(ctx context.Context, userId string) (*model.User, error) {
	db := repository.New(repository.Db)
	user, err := db.FetchUserByID(ctx, userId)
	if err != nil {
		log.Fatalf("failed to fetch user by name: %v", err)
	}

	return convertUser(user), nil
}
