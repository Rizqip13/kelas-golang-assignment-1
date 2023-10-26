package database

import (
	"assignment1/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	err    error
	dbPath = "score_assignment.db"
)

func StartDB() {
	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	err = db.AutoMigrate(models.Order{}, models.Item{})
	if err != nil {
		return
	}

	// di sqlite3 untuk enable ON DELETE CASCADE
	if err = db.Exec("PRAGMA foreign_keys = ON", nil).Error; err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
