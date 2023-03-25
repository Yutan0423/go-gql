package usecases

import "context"

type UserService interface {
	GetUserByName(ctx context.Context, name string)
}
