package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	h "students/pkg/http"
	"students/pkg/storage"
)

type App struct {
	Router     *chi.Mux
	Controller h.Controller
}

func (a *App) Initialize() {
	a.Router = chi.NewRouter()

	strg := storage.NewStorage()
	a.Controller = h.Controller{Storage: strg}

	a.Router.Use(middleware.RequestID)
	a.Router.Use(middleware.Logger)
	a.Router.Use(middleware.Recoverer)
	a.Router.Use(middleware.URLFormat)

	a.Router.Post("/create", a.Controller.UserCreate)
	a.Router.Post("/make_friends", a.Controller.UserAttach)
	a.Router.Delete("/user", a.Controller.UserDelete)
	a.Router.Get("/friends/{id}", a.Controller.UserGetFriends)
	a.Router.Put("/{id}", a.Controller.UserUpdate)
}

func (a *App) Run(httpPort string) {
	err := http.ListenAndServe(fmt.Sprintf(":%s", httpPort), a.Router)
	if err != nil {
		log.Fatal(err)
	}
}
