package user

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"refactoring/internal/app/domain/user"
	"time"
)

const(
	api = "api"
	apiVersion = "v1"
	urlUser = "/users"
	endPoint = "/"
)

type Handler struct {
	Service user.Service
}

func (h *Handler) Register(r chi.Router)  {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(time.Now().String()))
	})

	r.Route(api, func(r chi.Router) {
		r.Route(apiVersion, func(r chi.Router) {
			r.Route(urlUser, func(r chi.Router) {
				r.Get(endPoint, )
				r.Post(endPoint, )

				r.Route("/{id}", func(r chi.Router) {
					r.Get(endPoint, )
					r.Patch(endPoint, )
					r.Delete(endPoint, )
				})
			})
		})
	})
}
