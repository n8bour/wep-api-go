package main

import (
	"encoding/json"
	"fmt"
	"from_scratch_wep_api/models"
	"from_scratch_wep_api/pkg"
	"github.com/go-chi/chi/v5"
	"github.com/thedevsaddam/renderer"
	"log"
	"net/http"
)

type TestHandler struct {
	svc *pkg.TestService
}

func NewTestHandler(svc *pkg.TestService) *TestHandler {
	return &TestHandler{
		svc: svc,
	}
}

func (t *TestHandler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", t.PostTest)
	r.Get("/", t.GetTest)

	return r
}

func (t *TestHandler) GetTest(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set(renderer.ContentType, renderer.ContentJSON)

	test, err := t.svc.GetTest()
	if err != nil {
		log.Fatal(" error fetching test")
	}

	err = json.NewEncoder(writer).Encode(test)
	if err != nil {
		log.Fatal("error encoding...", err)
	}
}

func (t *TestHandler) PostTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(renderer.ContentType, renderer.ContentJSON)

	rs := models.TestRequest{}
	err := json.NewDecoder(r.Body).Decode(&rs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rs.Name)

	test, err := t.svc.CreateTest(rs)
	if err != nil {
		log.Fatal(err)
	}

	err = json.NewEncoder(w).Encode(&test)
	if err != nil {
		log.Fatal(err)
	}
}