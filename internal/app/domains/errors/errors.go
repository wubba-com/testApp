package domain

import (
	"errors"
)

var (
	// ErrUserNotFound user not found
	ErrUserNotFound = errors.New("user_not_found")
	// ErrEmpty field not empty
	ErrEmpty = errors.New("field must not be empty")
	// ErrEmail phone is not valid
	ErrEmail = errors.New("field email is not valid")
)