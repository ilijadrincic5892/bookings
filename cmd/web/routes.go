package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/ilijadrincic5892/bookings/pkg/config"
	"github.com/ilijadrincic5892/bookings/pkg/handlers"
)

// routes sets up the application routes
func routes(app *config.AppConfig) http.Handler {
	_ = app

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
