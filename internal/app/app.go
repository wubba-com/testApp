package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/wubba-com/testApp/internal/app/composite"
	"log"
	"net/http"
	"time"
)

func Run()  {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	//r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	u := composite.NewUserComposite(5 * time.Second)
	u.Handler.Register(r)

	log.Fatal(http.ListenAndServe(":3333", r))

}
