package config

import (
	"os"

	"github.com/me/goopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)	

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Errorf("file database not found: %v\n", err)
		
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

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite open error: %v\n", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite automigrate error: %v\n", err)
		return nil, err
	}

	return db, nil
}
