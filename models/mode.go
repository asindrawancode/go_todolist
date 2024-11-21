package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}

type ToDo struct {
	gorm.Model
	UserID uint
	Task   string
	Done   bool
}

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("todolist.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Migrate the schema
	DB.AutoMigrate(&User{}, &ToDo{})
}
