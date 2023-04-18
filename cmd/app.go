package main

import (
	"fmt"
	"from_scratch_wep_api/models"
	"from_scratch_wep_api/pkg"
	"from_scratch_wep_api/pkg/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewApp() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// using auth /JWT

	cfg := config.GetConfig()
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=postgres port=%s",
		cfg.DBUsername, cfg.DBPassword, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.Test{})
	if err != nil {
		log.Fatal(err)
	}

	serviceTest := pkg.NewServiceTest(db)

	r.Route("/api/v1", func(r chi.Router) {
		r.Mount("/test", serviceTest.Routes())
	})

	return r
}
