package models

import (
	"testing"

	m "github.com/XuryaX/card-dec-go/internal/models"
	"github.com/XuryaX/card-dec-go/internal/models/utils"
)

func TestNewCard(t *testing.T) {
	card := utils.NewCard("ACE", "SPADES")

	if card.Value != "ACE" {
		t.Errorf("Expected 'ACE', got '%s'", card.Value)
	}

	if card.Suit != "SPADES" {
		t.Errorf("Expected 'SPADES', got '%s'", card.Suit)
	}

	if card.Code != "AS" {
		t.Errorf("Expected 'AS', got '%s'", card.Code)
	}
}

func TestIsValidSuit(t *testing.T) {
	if !utils.IsValidSuit("SPADES") {
		t.Errorf("Expected 'SPADES' to be valid, but it was not")
	}

	if utils.IsValidSuit("INVALID_SUIT") {
		t.Errorf("Expected 'INVALID_SUIT' to be invalid, but it was valid")
	}
}

func TestIsValidValue(t *testing.T) {
	if !utils.IsValidValue("ACE") {
		t.Errorf("Expected 'ACE' to be valid, but it was not")
	}

	if utils.IsValidValue("INVALID_VALUE") {
		t.Errorf("Expected 'INVALID_VALUE' to be invalid, but it was valid")
	}
}

func TestIsValidCard(t *testing.T) {
	card := utils.NewCard("ACE", "SPADES")

	if !utils.IsValidCard(card) {
		t.Errorf("Expected card to be valid, but it was not")
	}

	invalidCard := m.Card{
		Value: "INVALID_VALUE",
		Suit:  "INVALID_SUIT",
		Code:  "IV",
	}

	if utils.IsValidCard(invalidCard) {
		t.Errorf("Expected invalid card to be invalid, but it was valid")
	}
}
