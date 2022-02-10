package composite

import (
	"github.com/wubba-com/testApp/internal/app/domain/user"
	"github.com/wubba-com/testApp/internal/app/user/delivery"
	user2 "github.com/wubba-com/testApp/internal/app/user/delivery/http/user"
	"github.com/wubba-com/testApp/internal/app/user/repository/jsonStorage"
	"github.com/wubba-com/testApp/internal/app/user/usecase"
	"time"
)

func NewUserComposite(timeout time.Duration) *UserComposite {
	r := jsonStorage.NewUserRepository()
	s := usecase.NewUserService(r, timeout)
	h := user2.NewUserHandler(s)

	return &UserComposite{
		Repository: r,
		Service: s,
		Handler: h,
	}
}

type UserComposite struct {
	Repository user.Repository
	Service user.Service
	Handler delivery.Handler
}