package app

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/wubba-com/testApp/internal/app/config"
	delivery "github.com/wubba-com/testApp/internal/app/user/delivery/http"
	"github.com/wubba-com/testApp/internal/app/user/repository/jsonStorage"
	"github.com/wubba-com/testApp/internal/app/user/usecase"
	js "github.com/wubba-com/testApp/pkg/client/json"
	"log"
	"net/http"
	"os"
	"time"
)

const(
	timeout = 3
)

func Run()  {
	// Получаем роутер
	r := chi.NewRouter()
	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Определяем middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	//r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// Получаем клиента взаимодействующий с хранилищем
	client := js.NewJsonClient(os.Getenv("STORAGE"))

	// Вызываем сервис
	repo := jsonStorage.NewUserRepository(client)
	s := usecase.NewUserService(repo, timeout * time.Second)
	h := delivery.NewUserHandler(s)
	h.Register(r)

	// Определяем порт из .env
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	fmt.Printf("start http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, r))

}
