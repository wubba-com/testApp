package user

import "net/http"

func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{}
}

type CreateUserRequest struct {
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

func (c *CreateUserRequest) Bind(r *http.Request) error { return nil }

func NewUpdateUserRequest() *UpdateUserRequest {
	return &UpdateUserRequest{}
}

type UpdateUserRequest struct {
	DisplayName string `json:"display_name"`
}

func (c *UpdateUserRequest) Bind(r *http.Request) error { return nil }