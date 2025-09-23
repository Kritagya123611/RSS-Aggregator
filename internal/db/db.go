package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error

	// Use modernc.org/sqlite instead of cgo sqlite
	DB, err = gorm.Open(sqlite.Open("file:app.db?cache=shared&mode=rwc&_pragma=foreign_keys(1)"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	log.Println("Connected to SQLite database (pure Go driver)!")
}
