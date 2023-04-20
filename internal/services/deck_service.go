package services

import (
	"errors"
	"fmt"

	"github.com/XuryaX/card-dec-go/internal/dal"
	e "github.com/XuryaX/card-dec-go/internal/exceptions"
	"github.com/XuryaX/card-dec-go/internal/models"
	"github.com/XuryaX/card-dec-go/internal/models/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// DeckService provides methods for managing decks.
type DeckService interface {
	CreateDeck(c *gin.Context) (*models.Deck, error)
	OpenDeck(c *gin.Context) (*models.Deck, error)
}

type DeckServiceImpl struct{}

func NewDeckService() DeckService {
	return &DeckServiceImpl{}
}

func (d *DeckServiceImpl) CreateDeck(c *gin.Context) (*models.Deck, error) {

	// NOTE:
	// Create Requests should always be POST.
	// Using Query Parameters instead of Request Body, since
	// specified in Req -- ?cards=AS,KD,AC,2C,KH.
	// Might be used for automated testing of correctness.

	// Get the "shuffled" and "cards" query parameters from the request
	shuffled := c.Query("shuffled") == "true"
	cardsParam := c.Query("cards")

	// Get the DeckDAL instance from the Gin context
	deckDAL := c.MustGet("deckDAL").(dal.DeckDAL)

	// Save the new deck in the database
	deck, err := deckDAL.CreateDeck(utils.NewDeck(shuffled, cardsParam))
	if err != nil {
		return nil, fmt.Errorf("failed to create deck: %v", err)
	}

	return deck, nil
}

func (d *DeckServiceImpl) OpenDeck(c *gin.Context) (*models.Deck, error) {

	// Get the deck ID from the query parameters
	deckID := c.Param("id")

	// Get the DeckDAL instance from the Gin context
	deckDAL := c.MustGet("deckDAL").(dal.DeckDAL)

	// Fetch the deck from the DAL
	deck, err := deckDAL.GetDeck(deckID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.ErrDeckNotFound
		}
		return nil, err
	}

	return deck, err
}
