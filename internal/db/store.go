package db

import "from_scratch_wep_api/models"

type Storer interface {
	GetTest() ([]*models.TestRequest, error)
	CreateTest(test models.TestRequest) (models.TestRequest, error)
}
