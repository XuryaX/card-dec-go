package main

import (
	"github.com/XuryaX/card-dec-go/api/handlers"
	"github.com/XuryaX/card-dec-go/internal/dal"
	"github.com/XuryaX/card-dec-go/internal/settings"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()

	dbConfig := settings.DatabaseConfig{
		Driver:     "sqlite",
		DSN:        "decks.db",
		GormConfig: &gorm.Config{},
	}

	database := dal.SetupDatabase(dbConfig)
	defer func() {
		sqlDB, _ := database.DB()
		sqlDB.Close()
	}()

	// Pass the DAL to the handlers
	handlers.InitHandlers(database)

	// Define API routes
	r.POST("/deck/new", handlers.CreateDeckHandler)
	r.GET("/deck/:id", handlers.OpenDeckHandler)
	r.GET("/deck/:id/draw", handlers.DrawCardHandler)

	// Start the Gin server
	r.Run(":8090")
}
