package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
	h "students/pkg/http"
	"students/pkg/storage"
)

func main() {
	r := chi.NewRouter()
	strg := storage.NewStorage()

	c := h.Controller{Storage: strg}

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	r.Post("/create", c.UserCreate)
	r.Post("/make_friends", c.UserAttach)
	r.Delete("/user", c.UserDelete)
	r.Get("/friends/{id}", c.UserGetFriends)
	r.Put("/{id}", c.UserUpdate)

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8081"
	}

	fmt.Println("Run port:", httpPort)

	err := http.ListenAndServe(fmt.Sprintf(":%s", httpPort), r)
	if err != nil {
		log.Fatal(err)
	}
}
