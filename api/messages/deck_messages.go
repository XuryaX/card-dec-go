package messages

import (
	m "github.com/XuryaX/card-dec-go/internal/models"
)

type DeckResponse struct {
	DeckID    string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}

type OpenDeckResponse struct {
	DeckID    string  `json:"deck_id"`
	Shuffled  bool    `json:"shuffled"`
	Remaining int     `json:"remaining"`
	Cards     m.Cards `json:"cards"`
}

type DrawCardResponse struct {
	Cards m.Cards `json:"cards"`
}
