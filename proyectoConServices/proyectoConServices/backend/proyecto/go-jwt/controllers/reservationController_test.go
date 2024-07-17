package controllers

import (
	"bytes"
	//"encoding/json"
	//"errors"
	//"proyecto/models"
	//"proyecto/services"

	//"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	//"time"

	//"proyecto/dtos"
	//"proyecto/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/mock"
)

// Returns a bad request status when input JSON is invalid
func TestCreateReservation_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/reservations", CreateReservation)

	invalidJSON := `{"HotelID": "invalid", "UserID": 1, "CheckIn": "2023-10-01", "CheckOut": "2023-10-10"}`

	req, _ := http.NewRequest(http.MethodPost, "/reservations", bytes.NewBufferString(invalidJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

//__________________________________________
/*
    // User is authenticated and has reservations
	func TestGetUserReservationsAuthenticated(t *testing.T) {
		// Setup
		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.GET("/reservations", GetUserReservations)

		// Mock user and reservations
		user := models.User{ID: 1}
		reservations := []models.Reservation{{ID: 1, UserID: 1, Date: time.Now()}}
		services.GetUserReservations = func(userID int) ([]models.Reservation, error) {
			return reservations, nil
		}

		// Create a request with the user in context
		req, _ := http.NewRequest(http.MethodGet, "/reservations", nil)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		c.Set("user", user)

		// Perform the request
		router.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusOK, w.Code)
		var response map[string][]models.Reservation
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, reservations, response["reservations"])
	}

*/
