package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.26

import (
	"context"
	"fmt"
	"go-gql/graphql/model"
	"go-gql/services"
)

// AddProjectV2ItemByID is the resolver for the addProjectV2ItemById field.
func (r *mutationResolver) AddProjectV2ItemByID(ctx context.Context, input model.AddProjectV2ItemByIDInput) (*model.AddProjectV2ItemByIDPayload, error) {
	panic(fmt.Errorf("not implemented: AddProjectV2ItemByID - addProjectV2ItemById"))
}

// Repository is the resolver for the repository field.
func (r *queryResolver) Repository(ctx context.Context, name string, owner string) (*model.Repository, error) {
	user, err := services.GetUserByName(ctx, owner)
	if err != nil {
		return nil, err
	}

	return services.GetRepoByFullName(ctx, user.ID, name)
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, name string) (*model.User, error) {
	return services.GetUserByName(ctx, name)
}

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	panic(fmt.Errorf("not implemented: Node - node"))
}

// Owner is the resolver for the owner field.
func (r *repositoryResolver) Owner(ctx context.Context, obj *model.Repository) (*model.User, error) {
	return services.GetUserByID(ctx, obj.Owner.ID)
}

// Issue is the resolver for the issue field.
func (r *repositoryResolver) Issue(ctx context.Context, obj *model.Repository, number int) (*model.Issue, error) {
	// とあるレポジトリに属する、とある番号のIssue情報を取得
	return services.GetIssueByRepoAndNumber(ctx, obj.ID, number)
}

// Issues is the resolver for the issues field.
func (r *repositoryResolver) Issues(ctx context.Context, obj *model.Repository, after *string, before *string, first *int, last *int) (*model.IssueConnection, error) {
	panic(fmt.Errorf("not implemented: Issues - issues"))
}

// PullRequest is the resolver for the pullRequest field.
func (r *repositoryResolver) PullRequest(ctx context.Context, obj *model.Repository, number int) (*model.PullRequest, error) {
	panic(fmt.Errorf("not implemented: PullRequest - pullRequest"))
}

// PullRequests is the resolver for the pullRequests field.
func (r *repositoryResolver) PullRequests(ctx context.Context, obj *model.Repository, after *string, before *string, first *int, last *int) (*model.PullRequestConnection, error) {
	panic(fmt.Errorf("not implemented: PullRequests - pullRequests"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Repository returns RepositoryResolver implementation.
func (r *Resolver) Repository() RepositoryResolver { return &repositoryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type repositoryResolver struct{ *Resolver }
