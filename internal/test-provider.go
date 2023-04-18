package internal

import (
	"from_scratch_wep_api/models"
	"gorm.io/gorm"
)

type TestFetcher interface {
	GetTest() (models.Test, error)
}

func NewTestFetcher(db *gorm.DB) TestRepository {
	return TestRepository{DB: db}
}

type TestRepository struct {
	DB *gorm.DB
}

func (t *TestRepository) GetTest() ([]models.Test, error) {
	var dest []models.Test
	t.DB.Model(&models.Test{}).Find(&dest)

	return dest, nil
}

func (t *TestRepository) CreateTest(test models.Test) (models.Test, error) {
	t.DB.Create(&test)

	return test, nil
}
