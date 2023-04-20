package handlers

import (
	"net/http"

	"github.com/XuryaX/card-dec-go/internal/dal"
	"github.com/gin-gonic/gin"
)

// CreateDeckHandler creates a new deck based on the provided options and returns it in JSON format.
func CreateDeckHandler(c *gin.Context) {
	// Get the "shuffled" and "cards" query parameters from the request
	shuffled := c.Query("shuffled") == "true"
	cardsParam := c.Query("cards")

	// Get the DeckDAL instance from the Gin context
	deckDAL := c.MustGet("deckDAL").(dal.DeckDAL)

	// Create a new deck based on the provided options
	deck, err := deckDAL.CreateDeck(shuffled, cardsParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return the created deck in JSON format
	c.JSON(http.StatusOK, deck)
}
