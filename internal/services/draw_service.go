package services

import (
	"errors"
	"strconv"

	"github.com/XuryaX/card-dec-go/internal/dal"
	e "github.com/XuryaX/card-dec-go/internal/exceptions"
	"github.com/XuryaX/card-dec-go/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// DeckService provides methods for managing decks.
type DrawService interface {
	DrawCards(c *gin.Context) ([]models.Card, error)
}

type DrawServiceImpl struct{}

func NewDrawService() DrawService {
	return &DrawServiceImpl{}
}

func (d *DrawServiceImpl) DrawCards(c *gin.Context) ([]models.Card, error) {

	// Get the DeckDAL instance from the Gin context
	deckDAL := c.MustGet("deckDAL").(dal.DeckDAL)

	// Get the deck ID from the request URL
	deckID := c.Param("id")

	// Get the count of cards to draw from the query parameters
	count, err := strconv.Atoi(c.Query("count"))
	if err != nil {
		return nil, e.InvalidCount
	}

	// Draw the requested cards from the deck
	cards, err := deckDAL.DrawCards(deckID, count)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.ErrDeckNotFound
		}
		return nil, err
	}

	return cards, nil
}
