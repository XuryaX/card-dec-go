package models

type Deck struct {
	ID        string `gorm:"primaryKey"`
	Shuffled  bool
	Remaining int
	Cards     []Card `gorm:"-"`
}

type DeckResponse struct {
	DeckID    string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}

type OpenDeckResponse struct {
	DeckID    string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
	Cards     []Card `json:"cards"`
}

type DrawCardResponse struct {
	Cards []Card `json:"cards"`
}
