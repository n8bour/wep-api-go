package pkg

import (
	"encoding/json"
	"fmt"
	"from_scratch_wep_api/internal"
	"from_scratch_wep_api/models"
	"github.com/go-chi/chi/v5"
	"github.com/thedevsaddam/renderer"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type TestService struct {
	TestRepository internal.TestRepository
}

func NewServiceTest(db *gorm.DB) TestService {
	return TestService{TestRepository: internal.NewTestRepository(db)}
}
func (t TestService) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", t.PostTest)
	r.Get("/", t.GetTest)

	return r
}

func (t TestService) GetTest(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set(renderer.ContentType, renderer.ContentJSON)

	test, err := t.TestRepository.GetTest()
	if err != nil {
		log.Fatal(" error fetching test")
	}

	err = json.NewEncoder(writer).Encode(test)
	if err != nil {
		log.Fatal("error encoding...", err)
	}
}

func (t TestService) PostTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(renderer.ContentType, renderer.ContentJSON)

	rs := models.Test{}
	err := json.NewDecoder(r.Body).Decode(&rs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rs.Name)

	test, err := t.TestRepository.CreateTest(rs)
	if err != nil {
		log.Fatal(err)
	}

	err = json.NewEncoder(w).Encode(&test)
	if err != nil {
		log.Fatal(err)
	}
}
