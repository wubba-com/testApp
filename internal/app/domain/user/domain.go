package user

import "time"

type (
	User struct {
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
