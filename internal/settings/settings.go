package settings

import "gorm.io/gorm"

// DatabaseConfig struct to store database configuration
type DatabaseConfig struct {
	Driver     string
	DSN        string
	GormConfig *gorm.Config
}

func GetDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Driver:     "sqlite",
		DSN:        "decks.db",
		GormConfig: &gorm.Config{},
	}
}
