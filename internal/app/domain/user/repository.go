package user

import "context"

type Repository interface {
	Get(context.Context, int) (*User, error)
	Store(context.Context, *CreateUserRequest) (int, error)
	Update(context.Context, *UpdateUserRequest) error
	Delete(context.Context, int) (int, error)
}