package main

import (
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
	// using auth /JWT

	routeHandler.Route("/api/v1", func(r chi.Router) {
		var (
			pdb         = db.NewPostgresProvider()
			svc         = pkg.NewTestService(pdb)
			testHandler = NewTestHandler(svc)
		)
		routeHandler.Mount("/test", testHandler.Routes())
	})

	err := http.ListenAndServe(":8080", routeHandler)
	if err != nil {
		log.Fatal("error in the server", err)
	}
}
