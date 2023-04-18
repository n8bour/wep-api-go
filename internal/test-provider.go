package internal

import "from_scratch_wep_api/models"

type TestFetcher interface {
	GetTest() (*models.Test, error)
}

func NewTestFetcher() TestRepository {
	return TestRepository{}
}

type TestRepository struct {
	// this will need a DB connection
}

func (t *TestRepository) GetTest() (*models.Test, error) {
	// this will fetch the test from the DB
	return &models.Test{Quantity: 1, Name: "Test"}, nil
}
