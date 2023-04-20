package dal

import (
	"errors"
	"fmt"

	e "github.com/XuryaX/card-dec-go/internal/exceptions"
	"github.com/XuryaX/card-dec-go/internal/models"
	"gorm.io/gorm"
)

type DeckDAL interface {
	CreateDeck(*models.Deck) (*models.Deck, error)
	GetDeck(deckID string) (*models.Deck, error)
	DrawCards(deckID string, count int) ([]models.Card, error)
}

type SQLiteDeckDAL struct {
	db *gorm.DB
}

func NewSQLiteDeckDAL(db *gorm.DB) DeckDAL {
	return &SQLiteDeckDAL{db: db}
}

func (d *SQLiteDeckDAL) CreateDeck(deck *models.Deck) (*models.Deck, error) {
	// Save the new deck in the database
	err := d.db.Create(deck).Error
	if err != nil {
		return nil, fmt.Errorf("failed to create deck: %v", err)
	}

	return deck, nil
}

func (d *SQLiteDeckDAL) GetDeck(deckID string) (*models.Deck, error) {
	var deck models.Deck

	// Retrieve the deck from the database using its ID
	err := d.db.First(&deck, "id = ?", deckID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("deck not found")
		}
		return nil, fmt.Errorf("failed to retrieve deck: %v", err)
	}

	return &deck, nil
}

func (d *SQLiteDeckDAL) DrawCards(deckID string, count int) ([]models.Card, error) {
	deck, err := d.GetDeck(deckID)
	if err != nil {
		return nil, err
	}

	if count > len(deck.Cards) {
		return nil, e.ErrInsufficientCards
	}

	drawnCards := deck.Cards[:count]
	deck.Cards = deck.Cards[count:]

	// Update the deck in the database
	err = d.db.Save(deck).Error
	if err != nil {
		return nil, fmt.Errorf("failed to update deck: %v", err)
	}

	return drawnCards, nil
}
