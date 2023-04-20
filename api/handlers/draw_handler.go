package handlers

import (
	"errors"
	"net/http"

	m "github.com/XuryaX/card-dec-go/api/messages"
	e "github.com/XuryaX/card-dec-go/internal/exceptions"
	s "github.com/XuryaX/card-dec-go/internal/services"
	"github.com/gin-gonic/gin"
)

// DrawCardHandler handles requests to draw one or more cards from a deck.
func DrawCardHandler(c *gin.Context) {

	drawService := c.MustGet("drawService").(s.DrawService)
	cards, err := drawService.DrawCards(c)

	if err != nil {

		if errors.Is(err, e.InvalidCount) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid count parameter",
			})
			return
		} else if errors.Is(err, e.ErrDeckNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Deck not found",
			})
			return
		} else if errors.Is(err, e.ErrInsufficientCards) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Insufficient cards in deck",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to draw cards from deck",
		})
		return
	}

	// Return the drawn cards in the response
	c.JSON(http.StatusOK, m.DrawCardResponse{
		Cards: cards,
	})
}
