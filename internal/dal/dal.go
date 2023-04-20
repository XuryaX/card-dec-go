package dal

import (
	"github.com/XuryaX/card-dec-go/internal/models"
	"github.com/XuryaX/card-dec-go/internal/settings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDatabase(config settings.DatabaseConfig) *gorm.DB {
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

	// Migrate the schema
	db.AutoMigrate(&models.Deck{})

	return db
}
