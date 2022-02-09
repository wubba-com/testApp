package user

import "context"

type Service interface {
	GetUser(int) (*User, error)
	StoreUser(*CreateUserRequest) (int, error)
	Update(*UpdateUserRequest) error
	Delete(int) (int, error)
}

func NewService(r Repository) Service {
	return &service{r}
}

type service struct {
	Repository Repository
}

func (s *service) GetUser(id int) (*User, error) {
	u, err := s.Repository.Get(context.Background(), id)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *service) StoreUser(input *CreateUserRequest) (int, error) {
	id, err := s.Repository.Store(context.Background(), input)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *service) Update(input *UpdateUserRequest) error {
	err := s.Repository.Update(context.Background(), input)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Delete(id int) (int, error) {
	id, err := s.Repository.Delete(context.Background(), id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

