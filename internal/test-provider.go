package internal

import (
	"from_scratch_wep_api/models"
	"gorm.io/gorm"
)

type TestDAO interface {
	GetTest() ([]models.Test, error)
	CreateTest(test models.Test) (models.Test, error)
}

func NewTestRepository(db *gorm.DB) TestRepository {
	return TestRepository{DB: db}
}

type TestRepository struct {
	DB *gorm.DB
}

func (t TestRepository) GetTest() (dest []models.Test, err error) {
	t.DB.Model(&models.Test{}).Find(&dest)

	return dest, nil
}

func (t TestRepository) CreateTest(test models.Test) (models.Test, error) {
	t.DB.Create(&test)

	return test, nil
}
