package usecases

import "context"

type userService interface {
	GetUserByName(ctx context.Context, name string)
}
