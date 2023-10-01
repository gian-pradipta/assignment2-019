package database

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

const DB_NAME string = "test.db"

func New() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(DB_NAME), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	db.Exec("PRAGMA foreign_keys = ON;")
	return db
}
