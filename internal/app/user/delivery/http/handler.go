package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/wubba-com/testApp/internal/app/user/delivery"
	"github.com/wubba-com/testApp/internal/app/user/delivery/http/validator"
	"github.com/wubba-com/testApp/pkg/render"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/wubba-com/testApp/internal/app/domains/user"

)

const(
	urlUser = "/users"
	endPoint = "/"
)

func NewUserHandler(s user.Service) delivery.Handler {
	return &Handler{s}
}

type Handler struct {
	Service user.Service
}

func (h *Handler) Register(r chi.Router)  {
	r.Get("/", h.welcome)

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route(urlUser, func(r chi.Router) {
				r.Get(endPoint, h.searchUsers)
				r.Post(endPoint, h.createUser)

				r.Route("/{id}", func(r chi.Router) {
					r.Get(endPoint, h.getUser)
					r.Patch(endPoint, h.updateUser)
					r.Delete(endPoint, h.deleteUser)
				})
			})
		})
	})
}

func (h *Handler) welcome(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte(time.Now().String()))
}

func (h *Handler) searchUsers(w http.ResponseWriter, r *http.Request)  {

	uStore, err := h.Service.SearchUsers(r.Context())
	if err != nil {
		log.Println(err)
		return
	}

	render.JSON(w, r, uStore.List)
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request)  {
	input := user.NewCreateUserRequest()

	if err := render.Bind(r, input); err != nil {
		log.Println(err)
		err = render.Render(w, r, res.ErrInvalidRequest(err))
		if err != nil {
			log.Println(err)
		}
		return
	}

	if ok, err := isCreateUserReqValid(input); !ok {
		err = render.Render(w, r, res.ErrInvalidRequest(err))
		if err != nil {
			log.Println(err)
		}
		return
	}

	id, err := h.Service.StoreUser(r.Context(), input)
	if err != nil {
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{
		"user_id": id,
	})
}

func (h *Handler) getUser(w http.ResponseWriter, r *http.Request)  {
	queryID := chi.URLParam(r, "id")

	id, err := strconv.Atoi(queryID)
	if err != nil {
		err = render.Render(w, r, res.ErrInvalidRequest(err))
		if err != nil {
			log.Printf(err.Error())
		}
		return
	}

	u, err := h.Service.GetUser(r.Context(), id)
	if err != nil {
		err = render.Render(w, r, res.ErrInvalidRequest(err))
		if err != nil {
			log.Printf(err.Error())
		}
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, u)
}

func (h *Handler) updateUser(w http.ResponseWriter, r *http.Request)  {
	input := user.NewUpdateUserRequest()

	if err := render.Bind(r, input); err != nil {
		err = render.Render(w, r, res.ErrInvalidRequest(err))
		if err != nil {
			log.Printf(err.Error())
		}
		return
	}

	queryID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(queryID)
	if err != nil {
		return
	}
	input.ID = id
	if ok, err := isUpdateUserReqValid(input); !ok {
		err = render.Render(w, r, res.ErrInvalidRequest(err))
		if err != nil {
			log.Printf(err.Error())
		}
		return
	}
	err = h.Service.UpdateUser(r.Context(), input)
	if err != nil {
		err = render.Render(w, r, res.ErrInvalidRequest(err))
		if err != nil {
			log.Printf(err.Error())
		}
		return
	}

	render.Status(r, http.StatusNoContent)
}

func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request)  {
	queryID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(queryID)
	if err != nil {
		return
	}

	id, err = h.Service.DeleteUser(r.Context(), id)
	if err != nil {
		err = render.Render(w, r, res.ErrInvalidRequest(err))
		if err != nil {
			return
		}
	}

	render.Status(r, http.StatusNoContent)
}

func isCreateUserReqValid(input *user.CreateUserRequest) (bool, error) {
	v := validator.CreateUserRequestValid{}
	return v.IsValid(input)
}

func isUpdateUserReqValid(input *user.UpdateUserRequest) (bool, error) {
	v := validator.UpdateUserRequestValid{}
	return v.IsValid(input)
}