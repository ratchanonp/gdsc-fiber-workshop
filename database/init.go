package database

import (
	"fmt"
	"workshop/project/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func InitDB() {
	var err error
	DBConn, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Database connection successfully opened")

	DBConn.AutoMigrate(&model.Book{})
	fmt.Println("Database Migrated")
}
