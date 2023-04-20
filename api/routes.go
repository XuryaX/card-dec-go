package api

import (
	"github.com/XuryaX/card-dec-go/api/handlers"
	"github.com/gin-gonic/gin"
)

func Init_routes(r *gin.Engine) {
	// Define API routes
	r.POST("/deck/new", handlers.CreateDeckHandler)
	r.GET("/deck/:id", handlers.OpenDeckHandler)
	r.POST("/deck/:id/draw", handlers.DrawCardHandler)
}
