package main

import (
	"from_scratch_wep_api/pkg"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewApp() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// using auth /JWT

	r.Route("/api/v1", func(r chi.Router) {
		r.Mount("/test", pkg.TestService{}.Routes())
	})

	return r
}
