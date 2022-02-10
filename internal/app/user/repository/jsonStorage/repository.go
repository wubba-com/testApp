package jsonStorage

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/wubba-com/testApp/internal/app/domain/user"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)


const store = `users.json`
var (
	UserNotFound = errors.New("user_not_found")
)
func NewUserRepository() user.Repository {
	return &userRepository{}
}

type userRepository struct {

}

func (ur *userRepository) Get(ctx context.Context, id int) (*user.User, error) {
	dir, err := os.Getwd()
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	var storage = filepath.Join(dir, store)

	s := user.NewStore()
	f, err := ioutil.ReadFile(storage)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(f, s)
	if err != nil {
		return nil, err
	}
	key := strconv.Itoa(id)

	if u, ok := s.List[key]; ok {
		return &u, nil
	} else {
		return nil, UserNotFound
	}
}

func (ur *userRepository) All(ctx context.Context) (*user.Store, error) {
	f, err := ioutil.ReadFile(store)
	if err != nil {
		return nil, err
	}
	s := user.NewStore()
	err = json.Unmarshal(f, s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (ur *userRepository) Store(ctx context.Context, input *user.CreateUserRequest) (int, error) {
	f, _ := ioutil.ReadFile(store)
	s := user.NewStore()
	err := json.Unmarshal(f, &s)
	if err != nil {
		return 0, err
	}

	s.Increment++

	u := user.User{
		CreatedAt:   time.Now(),
		DisplayName: input.DisplayName,
		Email:       input.DisplayName,
	}

	id := strconv.Itoa(s.Increment)
	s.List[id] = u

	b, err := json.Marshal(&s)
	if err != nil {
		return 0, err
	}

	err = ioutil.WriteFile(store, b, fs.ModePerm)
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
	f, err := ioutil.ReadFile(store)
	if err != nil {
		return err
	}
	err = json.Unmarshal(f, s)
	if err != nil {
		return err
	}
	userID := strconv.Itoa(input.ID)

	if _, ok := s.List[userID]; !ok {
		return UserNotFound
	}

	u := s.List[userID]



	u.DisplayName = input.DisplayName
	u.Email = input.Email
	s.List[userID] = u

	b, err := json.Marshal(s)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(store, b, fs.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) Delete(ctx context.Context, id int) (int, error) {
	f, _ := ioutil.ReadFile(store)
	s := user.NewStore()
	_ = json.Unmarshal(f, &s)

	userID := strconv.Itoa(id)
	if _, ok := s.List[userID]; !ok {
		return 0, errors.New("error")
	}

	delete(s.List, userID)

	b, err := json.Marshal(&s)
	if err != nil {
		return 0, err
	}
	err = ioutil.WriteFile(store, b, fs.ModePerm)
	if err != nil {
		return 0, err
	}

	return id, nil
}

