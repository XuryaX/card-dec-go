package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/XuryaX/card-dec-go/api/handlers"
	m "github.com/XuryaX/card-dec-go/api/messages"
	"github.com/XuryaX/card-dec-go/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockDeckService struct {
	mock.Mock
}

func (m *mockDeckService) CreateDeck(c *gin.Context) (*models.Deck, error) {
	args := m.Called(c)
	return args.Get(0).(*models.Deck), args.Error(1)
}

func (m *mockDeckService) OpenDeck(c *gin.Context) (*models.Deck, error) {
	args := m.Called(c)
	return args.Get(0).(*models.Deck), args.Error(1)
}

func TestCreateDeckHandler(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Mock DeckService
	deckService := new(mockDeckService)
	deckService.On("CreateDeck", c).Return(&models.Deck{
		ID:       "test-deck",
		Shuffled: true,
		Cards:    []models.Card{},
	}, nil)

	c.Set("deckService", deckService)

	// Execute
	handlers.CreateDeckHandler(c)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	var response m.DeckResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "test-deck", response.DeckID)
	assert.Equal(t, true, response.Shuffled)
	assert.Equal(t, 0, response.Remaining)
}

func TestOpenDeckHandler(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Mock DeckService
	deckService := new(mockDeckService)
	deckService.On("OpenDeck", c).Return(&models.Deck{
		ID:       "test-deck",
		Shuffled: true,
		Cards:    []models.Card{},
	}, nil)

	c.Set("deckService", deckService)

	// Execute
	handlers.OpenDeckHandler(c)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	var response m.OpenDeckResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "test-deck", response.DeckID)
	assert.Equal(t, true, response.Shuffled)
	assert.Equal(t, 0, response.Remaining)
	assert.Equal(t, models.Cards{}, response.Cards)
}
