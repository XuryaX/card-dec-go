package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	h "github.com/XuryaX/card-dec-go/api/handlers"
	e "github.com/XuryaX/card-dec-go/internal/exceptions"
	m "github.com/XuryaX/card-dec-go/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDrawService struct {
	mock.Mock
}

func (mds *MockDrawService) DrawCards(c *gin.Context) ([]m.Card, error) {
	args := mds.Called(c)
	return args.Get(0).([]m.Card), args.Error(1)
}

func TestDrawCardHandler_Success(t *testing.T) {
	// Create a mock draw service
	mockDrawService := new(MockDrawService)
	mockCards := []m.Card{{Value: "A", Suit: "Spades"}, {Value: "2", Suit: "Spades"}}
	mockDrawService.On("DrawCards", mock.AnythingOfType("*gin.Context")).Return(mockCards, nil)

	// Create a test context and set the drawService
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("drawService", mockDrawService)

	// Make a request to the handler
	req, _ := http.NewRequest("GET", "/decks/1/draw?count=2", nil)
	c.Request = req
	h.DrawCardHandler(c)

	// Check the response status code and JSON response
	assert.Equal(t, http.StatusOK, w.Code)

	expectedJSON := `{"cards":[{"value":"A","suit":"Spades","code":""},{"value":"2","suit":"Spades","code":""}]}`
	assert.Equal(t, expectedJSON, strings.TrimSpace(w.Body.String()))

	// Check that the mock draw service was called correctly
	mockDrawService.AssertExpectations(t)
}

func TestDrawCardHandler_InvalidCount(t *testing.T) {
	// Create a mock draw service
	mockDrawService := new(MockDrawService)

	mockDrawService.On("DrawCards", mock.AnythingOfType("*gin.Context")).Return([]m.Card{}, e.InvalidCount)

	// Create a test context and set the drawService
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("drawService", mockDrawService)

	// Make a request to the handler
	req, _ := http.NewRequest("GET", "/decks/1/draw?count=0", nil)
	c.Request = req
	h.DrawCardHandler(c)

	// Check the response status code and JSON response
	assert.Equal(t, http.StatusBadRequest, w.Code)

	expectedJSON := `{"error":"Invalid count parameter"}`
	assert.Equal(t, expectedJSON, strings.TrimSpace(w.Body.String()))

	// Check that the mock draw service was called correctly
	mockDrawService.AssertExpectations(t)
}
