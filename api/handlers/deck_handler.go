package handlers

import (
	"errors"
	"net/http"

	m "github.com/XuryaX/card-dec-go/api/messages"
	e "github.com/XuryaX/card-dec-go/internal/exceptions"
	s "github.com/XuryaX/card-dec-go/internal/services"
	"github.com/gin-gonic/gin"
)

// CreateDeckHandler creates a new deck based on the provided options and returns it in JSON format.
func CreateDeckHandler(c *gin.Context) {

	deckService := c.MustGet("deckService").(s.DeckService)
	// Create a new deck based on the provided options
	deck, err := deckService.CreateDeck(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Build the response JSON
	response := m.DeckResponse{
		DeckID:    deck.ID,
		Shuffled:  deck.Shuffled,
		Remaining: len(deck.Cards)}

	// Return the created deck in JSON format
	c.JSON(http.StatusOK, response)
}

// OpenDeckHandler is an HTTP handler that returns the details of an open deck.
func OpenDeckHandler(c *gin.Context) {

	deckService := c.MustGet("deckService").(s.DeckService)

	deck, err := deckService.OpenDeck(c)
	if err != nil {
		if errors.Is(err, e.ErrDeckNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to fetch remaining cards for deck",
			})
			return
		} else {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
			return
		}
	}

	// Build the response JSON
	response := m.OpenDeckResponse{
		DeckID:    deck.ID,
		Shuffled:  deck.Shuffled,
		Remaining: len(deck.Cards),
		Cards:     deck.Cards,
	}

	// Return the response JSON
	c.JSON(http.StatusOK, response)
}
