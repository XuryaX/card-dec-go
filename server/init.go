package server

import (
	"github.com/XuryaX/card-dec-go/internal/dal"
	"github.com/XuryaX/card-dec-go/internal/services"
	"github.com/XuryaX/card-dec-go/internal/settings"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Initializes Dependencies & does Dependency Injection into gin context

func initDependencies(r *gin.Engine) *gorm.DB {
	dbConfig := settings.GetDatabaseConfig()
	db := setupDB()

	switch dbConfig.Driver {
	case "sqlite":
		sqlitedb := dal.NewSQLiteDeckDAL(db)
		r.Use(func(c *gin.Context) {
			c.Set("deckDAL", sqlitedb)
			c.Next()
		})
	default:
		panic("unsupported database driver")
	}

	deckService := services.NewDeckService()
	r.Use(func(c *gin.Context) {
		c.Set("deckService", deckService)
		c.Next()
	})

	drawService := services.NewDrawService()

	r.Use(func(c *gin.Context) {
		c.Set("drawService", drawService)
		c.Next()
	})

	return db
}
