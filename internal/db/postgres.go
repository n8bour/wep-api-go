package db

import (
	"fmt"
	"from_scratch_wep_api/models"
	"from_scratch_wep_api/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Postgres struct {
	*gorm.DB
}

func NewPostgresProvider() *Postgres {
	cfg := config.GetConfig()
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=postgres port=%s",
		cfg.DBUsername, cfg.DBPassword, cfg.DBPort)

	pdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return &Postgres{DB: pdb}
}

func (p *Postgres) SeedDB(a ...any) error {
	return p.AutoMigrate(&a)
}

func (p *Postgres) GetTest() (dest []*models.TestRequest, err error) {
	p.Model(&models.TestRequest{}).Find(&dest)

	return dest, nil
}

func (p *Postgres) CreateTest(test models.TestRequest) (models.TestRequest, error) {
	p.Create(&test)

	return test, nil
}
