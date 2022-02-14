package jsonStorage

import (
	"context"
	domain "github.com/wubba-com/testApp/internal/app/domains/errors"
	"github.com/wubba-com/testApp/internal/app/domains/user"
	js "github.com/wubba-com/testApp/pkg/client/json"
	"strconv"
	"time"
)

func NewUserRepository(client *js.JsonStorage) user.Repository {
	return &userRepository{js: client}
}

type userRepository struct {
	js *js.JsonStorage
}

func (ur *userRepository) Get(ctx context.Context, id int) (*user.User, error) {
	s := user.NewStore()
	err := ur.js.Encode(s)
	if err != nil {
		return nil, err
	}

	key := strconv.Itoa(id)

	if u, ok := s.List[key]; ok {
		return &u, nil
	} else {
		return nil, domain.ErrUserNotFound
	}
}

func (ur *userRepository) All(ctx context.Context) (*user.Store, error) {
	s := user.NewStore()
	err := ur.js.Encode(s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (ur *userRepository) Store(ctx context.Context, input *user.CreateUserRequest) (int, error) {
	s := user.NewStore()
	err := ur.js.Encode(s)
	if err != nil {
		return 0, err
	}

	s.Increment++

	u := user.User{
		DisplayName: input.DisplayName,
		Email:       input.Email,
		CreatedAt:   time.Now(),
		UpdateAt: time.Now(),
	}

	id := strconv.Itoa(s.Increment)
	s.List[id] = u

	b, err := ur.js.Decode(s)
	if err != nil {
		return 0, err
	}

	_, err = ur.js.Write(b)
	if err != nil {
		return 0, err
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (ur *userRepository) Update(ctx context.Context, input *user.UpdateUserRequest) error {
	s := user.NewStore()
	err := ur.js.Encode(s)
	if err != nil {
		return err
	}
	userID := strconv.Itoa(input.ID)

	if _, ok := s.List[userID]; !ok {
		return domain.ErrUserNotFound
	}

	u := s.List[userID]

	u.DisplayName = input.DisplayName
	u.Email = input.Email
	u.UpdateAt = time.Now()
	s.List[userID] = u

	b, err := ur.js.Decode(s)
	if err != nil {
		return err
	}
	_, err = ur.js.Write(b)
	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) Delete(ctx context.Context, id int) (int, error) {
	s := user.NewStore()
	err := ur.js.Encode(s)

	userID := strconv.Itoa(id)
	if _, ok := s.List[userID]; !ok {
		return 0, domain.ErrUserNotFound
	}

	delete(s.List, userID)

	b, err := ur.js.Decode(s)
	if err != nil {
		return 0, err
	}
	_, err = ur.js.Write(b)
	if err != nil {
		return 0, err
	}

	return id, nil
}
