package models

import "gorm.io/gorm"

type Test struct {
	gorm.Model
	Name     string `json:"name"`
	Quantity uint   `json:"quantity"`
}
