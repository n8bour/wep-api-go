package pkg

import (
	"context"
	"encoding/json"
	"from_scratch_wep_api/config"
	"from_scratch_wep_api/internal"
	"from_scratch_wep_api/types"
	"github.com/go-chi/chi/v5"
	"github.com/thedevsaddam/renderer"
	"log"
	"net/http"
)

type TestHandler struct {
	svc *internal.TestService
}

func NewTestHandler(cfg *config.Config) *TestHandler {
	return &TestHandler{
		svc: internal.NewTestService(cfg),
	}
}

func (t *TestHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Route("/test", func(r chi.Router) {
		ctx := context.Background()
		r.Post("/", t.PostTest(ctx))
		r.Get("/", t.GetTests(ctx))
	})

	return r
}

func (t *TestHandler) GetTests(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(renderer.ContentType, renderer.ContentJSON)
		test, err := t.svc.GetTests(ctx)
		if err != nil {
			log.Fatal(" error fetching test")
		}

		err = json.NewEncoder(w).Encode(test)
		if err != nil {
			log.Fatal("error encoding...", err)
		}
	}
}

func (t *TestHandler) PostTest(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(renderer.ContentType, renderer.ContentJSON)

		tr := types.TestRequest{}
		err := json.NewDecoder(r.Body).Decode(&tr)
		if err != nil {
			log.Fatal(err)
		}

		test, err := t.svc.CreateTest(ctx, tr)
		if err != nil {
			log.Fatal(err)
		}

		err = json.NewEncoder(w).Encode(&test)
		if err != nil {
			log.Fatal(err)
		}
	}
}
