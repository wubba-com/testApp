package validator

import (
	domain "github.com/wubba-com/testApp/internal/app/domains/errors"
	"github.com/wubba-com/testApp/internal/app/domains/user"
	"strings"
)


func With(field string) func(string) int {
	return func(search string) int {
		return strings.Index(field, search)
	}
}

type CreateUserRequestValid struct {

}

func (crv *CreateUserRequestValid) IsValid(input *user.CreateUserRequest) (bool, error) {
	if input.DisplayName == "" {
		return false, domain.ErrEmpty
	}
	if input.Email == "" {
		return false, domain.ErrEmpty
	}
	index := With(input.Email)
	if index("@") < 0 {
		return false, domain.ErrEmail
	}
	return true, nil
}

type UpdateUserRequestValid struct {

}

func (crv *UpdateUserRequestValid) IsValid(input *user.UpdateUserRequest) (bool, error) {
	if input.DisplayName == "" {
		return false, domain.ErrEmpty
	}
	if input.Email == "" {
		return false, domain.ErrEmpty
	}
	index := With(input.Email)
	if index("@") < 0 {
		return false, domain.ErrEmail
	}
	return true, nil
}
