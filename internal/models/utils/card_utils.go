package utils

import (
	"github.com/XuryaX/card-dec-go/internal/models"
)

// NewCard creates a new Card with the specified value and suit.
func NewCard(value, suit string) models.Card {
	return models.Card{
		Value: value,
		Suit:  suit,
		Code:  value[:1] + suit[:1],
	}
}

// IsValidSuit checks if a given suit is valid.
func IsValidSuit(suit string) bool {
	validSuits := []string{"SPADES", "DIAMONDS", "CLUBS", "HEARTS"}

	for _, validSuit := range validSuits {
		if suit == validSuit {
			return true
		}
	}
	return false
}

// IsValidValue checks if a given card value is valid.
func IsValidValue(value string) bool {
	validValues := []string{"ACE", "2", "3", "4", "5", "6", "7", "8", "9", "10", "JACK", "QUEEN", "KING"}

	for _, validValue := range validValues {
		if value == validValue {
			return true
		}
	}
	return false
}

// IsValidCard checks if a given card is valid.
func IsValidCard(card models.Card) bool {
	return IsValidSuit(card.Suit) && IsValidValue(card.Value)
}
