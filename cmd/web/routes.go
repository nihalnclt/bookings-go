package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/nihalnclt/bookings-go/internal/handlers"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	mux.Get("/user/login", handlers.Repo.ShowLogin)
	mux.Post("/user/login", handlers.Repo.PostShowLogin)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
