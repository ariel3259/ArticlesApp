package database

import (
	"ArticlesApi/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetCon() *gorm.DB {
	dsn := "root:1234@tcp(localhost:3306)/articles_db?parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Printf(err.Error())
		return nil
	}
	return db
}

func MakeModels(db *gorm.DB) {
	db.AutoMigrate(model.Articles{})
}
