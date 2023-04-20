package server

import (
	"github.com/XuryaX/card-dec-go/internal/models"
	"github.com/XuryaX/card-dec-go/internal/settings"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupDB() *gorm.DB {
	db_settings := settings.GetDatabaseConfig()
	database := setupDatabase(db_settings)
	return database
}

func migrateDB() *gorm.DB {
	db_settings := settings.GetDatabaseConfig()
	database := setupDatabase(db_settings)

	// Migrate the schema
	database.AutoMigrate(&models.Deck{})

	return database
}

func setupDatabase(config settings.DatabaseConfig) *gorm.DB {
	var dialector gorm.Dialector

	switch config.Driver {
	case "sqlite":
		dialector = sqlite.Open(config.DSN)
	default:
		panic("unsupported database driver")
	}

	db, err := gorm.Open(dialector, config.GormConfig)
	if err != nil {
		panic("failed to connect to the database")
	}

	return db
}
