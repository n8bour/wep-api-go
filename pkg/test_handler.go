package pkg

import (
	"context"
	"encoding/json"
	"from_scratch_wep_api/internal"
	"from_scratch_wep_api/internal/db"
	"from_scratch_wep_api/types"
	"github.com/go-chi/chi/v5"
	"github.com/thedevsaddam/renderer"
	"log"
	"net/http"
)

type TestHandler struct {
	svc    *internal.TestService
	render *renderer.Render
}

func NewTestHandler(store db.Storer) *TestHandler {
	return &TestHandler{
		svc:    internal.NewTestService(store),
		render: &renderer.Render{},
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
		test, err := t.svc.GetTests(ctx)
		if err != nil {
			log.Fatal(" error fetching test")
		}

		err = t.render.JSON(w, http.StatusOK, test)
		if err != nil {
			log.Fatal("error encoding...", err)
		}
	}
}

func (t *TestHandler) PostTest(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tr := types.TestRequest{}
		err := json.NewDecoder(r.Body).Decode(&tr)
		if err != nil {
			_ = t.render.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		}

		test, err := t.svc.CreateTest(ctx, tr)
		if err != nil {
			_ = t.render.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		}

		err = t.render.JSON(w, http.StatusOK, test)
		if err != nil {
			log.Fatal(err)
		}
	}
}
