package tests

import (
	"cab-booking-system/models"
	"cab-booking-system/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTripRepository_CreateTrip(t *testing.T) {
	repo := repository.NewTripRepository()
	riderID := "r1"
	cabID := "cab1"
	source := models.Location{Latitude: 12.0, Longitude: 77.0}
	dest := models.Location{Latitude: 13.0, Longitude: 78.0}
	trip, err := repo.CreateTrip(riderID, cabID, source, dest)
	assert.NoError(t, err)
	assert.Equal(t, riderID, trip.RiderID)
	assert.Equal(t, cabID, trip.CabID)
}

func TestTripRepository_GetAllTripsByRider(t *testing.T) {
	repo := repository.NewTripRepository()
	riderID := "r1"
	cabID := "cab1"
	source := models.Location{Latitude: 12.0, Longitude: 77.0}
	dest := models.Location{Latitude: 13.0, Longitude: 78.0}
	_, _ = repo.CreateTrip(riderID, cabID, source, dest)
	trips, err := repo.GetAllTripsByRider(riderID)
	assert.NoError(t, err)
	assert.NotEmpty(t, trips)

	trips, err = repo.GetAllTripsByRider("notfound")
	assert.NoError(t, err)
	assert.Empty(t, trips)
}

func TestTripRepository_EndTrip(t *testing.T) {
	repo := repository.NewTripRepository()
	riderID := "r1"
	cabID := "cab1"
	source := models.Location{Latitude: 12.0, Longitude: 77.0}
	dest := models.Location{Latitude: 13.0, Longitude: 78.0}
	trip, _ := repo.CreateTrip(riderID, cabID, source, dest)

	err := repo.EndTrip(trip.ID)
	assert.NoError(t, err)
	assert.Equal(t, models.TripStatusFinished, trip.Status)
}
