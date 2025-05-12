package config

import (
	"os"

	"github.com/pierrialexandervidmar/go-jobs/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/data.db"

	// Check if the database exists
	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("Database file not found, creating...")

		// Create database file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}

		file.Close()
	}

	// Create DB e Connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("Sqlite opening error: %v", err)
		return nil, err
	}

	// Migrate Schema
	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil {
		logger.Errorf("Sqlite AutoMigration error: %v", err)
		return nil, err
	}

	// Return DB
	return db, nil
}
