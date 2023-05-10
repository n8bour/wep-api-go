package db

import (
	"fmt"
	"from_scratch_wep_api/config"
	"from_scratch_wep_api/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Postgres struct {
	*gorm.DB
}

func NewPostgresProvider(cfg *config.Config) *Postgres {
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=postgres port=%s",
		cfg.DBUsername, cfg.DBPassword, cfg.DBPort)

	pdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	//create tables
	err = pdb.AutoMigrate(&types.Test{})
	if err != nil {
		log.Fatal(err)
	}

	return &Postgres{DB: pdb}
}

func (p *Postgres) GetTests() (dest []*types.TestRequest, err error) {
	var tests []*types.Test

	p.Model(&types.Test{}).Find(&tests)

	for _, t := range tests {
		elem := t.ToTestRequest()
		dest = append(dest, &elem)
	}

	return dest, nil
}

func (p *Postgres) CreateTest(testReq types.TestRequest) (types.TestRequest, error) {
	test := testReq.ToTest()

	p.Create(&test)

	return test.ToTestRequest(), nil
}
