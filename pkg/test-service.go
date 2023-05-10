package pkg

import (
	"from_scratch_wep_api/internal/db"
	"from_scratch_wep_api/models"
)

type TestService struct {
	store db.Storer
}

func NewTestService(s db.Storer) *TestService {
	return &TestService{store: s}
}

func (t *TestService) GetTest() ([]*models.TestRequest, error) {
	return t.store.GetTest()
}

func (t *TestService) CreateTest(ts models.TestRequest) (models.TestRequest, error) {
	return t.store.CreateTest(ts)
}
