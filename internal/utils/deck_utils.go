package utils

import (
	"github.com/XuryaX/card-dec-go/internal/models"

	"github.com/google/uuid"
)

func GenerateDeckID() string {
	return uuid.New().String()
}

func GenerateDeck(shuffled bool, cards []string) models.Deck {
	// TODO: Implement deck generation logic
	return models.Deck{}
}

func DrawCards(deck *models.Deck, count int) []models.Card {
	// TODO: Implement card drawing logic
	return []models.Card{}
}
