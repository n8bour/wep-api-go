package main

import (
	"encoding/json"
	"from_scratch_wep_api/config"
	"from_scratch_wep_api/internal/db"
	"from_scratch_wep_api/pkg"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func main() {
	routeHandler := chi.NewRouter()
	routeHandler.Use(middleware.Logger)
	routeHandler.Use(middleware.Recoverer)

	var cfg = config.GetConfig()

	routeHandler.Group(func(r chi.Router) {
		// using auth /JWT && maybe change the way to pass postgres
		var testHandler = pkg.NewTestHandler(db.NewPostgresProvider(cfg))
		r.Mount("/api/v1", testHandler.Routes())
	})

	// public endpoint
	routeHandler.Group(func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
			_ = json.NewEncoder(w).Encode("Hello World")
		})
	})

	log.Printf("Server is up and running on port: %s", cfg.Addr)
	err := http.ListenAndServe(cfg.Addr, routeHandler)
	if err != nil {
		log.Fatal("error in the server", err)
	}
}
