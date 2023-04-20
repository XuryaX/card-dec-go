package utils

import (
	"math/rand"
	"strings"

	"github.com/XuryaX/card-dec-go/internal/models"
	"github.com/google/uuid"
)

// Suits represents the possible suits for a card in a deck.
var Suits = []string{"SPADES", "HEARTS", "DIAMONDS", "CLUBS"}

// Values represents the possible values for a card in a deck.
var Values = []string{"ACE", "2", "3", "4", "5", "6", "7", "8", "9", "10", "JACK", "QUEEN", "KING"}

// ParseCardCode parses a card code (e.g. "AS", "10H") and returns its value and suit.
func ParseCardCode(cardCode string) (value string, suit string) {
	if len(cardCode) != 2 {
		return "", ""
	}

	value = string(cardCode[0])
	suit = string(cardCode[1])

	if value == "1" && len(cardCode) == 3 {
		value = "10"
		suit = string(cardCode[2])
	}

	return value, suit
}

// NewDeck creates a new deck with the specified shuffling and cards parameters.
func NewDeck(shuffled bool, cardsParam string) *models.Deck {
	deck := &models.Deck{
		ID:        uuid.New().String(),
		Shuffled:  shuffled,
		Remaining: 0,
		Cards:     []models.Card{},
	}

	if cardsParam == "" {
		for _, suit := range Suits {
			for _, value := range Values {
				card := NewCard(value, suit)
				deck.Cards = append(deck.Cards, card)
			}
		}
	} else {
		cardStrings := strings.Split(cardsParam, ",")
		for _, cardString := range cardStrings {
			value, suit := ParseCardCode(cardString)
			if IsValidValue(value) && IsValidSuit(suit) {
				card := NewCard(value, suit)
				deck.Cards = append(deck.Cards, card)
			}
		}
	}

	deck.Remaining = len(deck.Cards)

	if shuffled {
		rand.Shuffle(len(deck.Cards), func(i, j int) {
			deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
		})
	}

	return deck
}
