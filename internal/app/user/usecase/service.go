package usecase

import (
	"context"
	"github.com/wubba-com/testApp/internal/app/domain/user"
	"time"
)

func NewUserService(r user.Repository, timeout time.Duration) user.Service {
	return &service{r, timeout}
}

type service struct {
	Repository user.Repository
	CtxTimeout time.Duration
}

func (s *service) SearchUsers(ctx context.Context) (*user.Store, error) {
	ctx, cancel := context.WithTimeout(ctx, s.CtxTimeout)
	defer cancel()

	u, err := s.Repository.All(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *service) GetUser(ctx context.Context ,id int) (*user.User, error) {
	ctx, cancel := context.WithTimeout(ctx, s.CtxTimeout)
	defer cancel()

	u, err := s.Repository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *service) StoreUser(ctx context.Context, input *user.CreateUserRequest) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, s.CtxTimeout)
	defer cancel()

	id, err := s.Repository.Store(ctx, input)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *service) UpdateUser(ctx context.Context, input *user.UpdateUserRequest) error {
	ctx, cancel := context.WithTimeout(ctx, s.CtxTimeout)
	defer cancel()

	err := s.Repository.Update(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteUser(ctx context.Context ,id int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, s.CtxTimeout)
	defer cancel()

	id, err := s.Repository.Delete(ctx, id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
