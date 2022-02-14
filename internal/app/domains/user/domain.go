package user

import (
	"context"
	"time"
)

func NewStore() *Store {
	return &Store{}
}

type (
	User struct {
		UpdateAt time.Time `json:"update_at"`
		CreatedAt   time.Time `json:"created_at"`
		DisplayName string    `json:"display_name"`
		Email       string    `json:"email"`
	}
	List map[string]User

	Store struct {
		Increment int  `json:"increment"`
		List      List `json:"list"`
	}
)

type Repository interface {
	Get(context.Context, int) (*User, error)
	All(context.Context) (*Store, error)
	Store(context.Context, *CreateUserRequest) (int, error)
	Update(context.Context, *UpdateUserRequest) error
	Delete(context.Context, int) (int, error)
}

type Service interface {
	SearchUsers(context.Context) (*Store, error)
	GetUser(context.Context, int) (*User, error)
	StoreUser(context.Context, *CreateUserRequest) (int, error)
	UpdateUser(context.Context, *UpdateUserRequest) error
	DeleteUser(context.Context, int) (int, error)
}