package model

import "gorm.io/gorm"

type Articles struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Stock       uint    `json:"stock"`
	State       bool
}
