package server

import (
	"github.com/XuryaX/card-dec-go/api"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()
	database := initDependencies(r)
	api.Init_routes(r)

	defer func() {
		sqlDB, _ := database.DB()
		sqlDB.Close()
	}()

	// Start the Gin server
	r.Run(":8090")
}
