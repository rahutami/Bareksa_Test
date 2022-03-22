package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "data.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&News{})
	database.AutoMigrate(&Tag{})
	database.AutoMigrate(&NewsTags{})

	DB = database
}