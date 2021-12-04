package database

import (
	"golang_architecture_ddd/domain/model"

	"github.com/jinzhu/gorm"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("mysql", "user:password@tcp(ddd_db)/ddd?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(model.Task{})

	return db
}
