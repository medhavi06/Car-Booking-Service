package tests

import (
	"cab-booking-system/models"
	"cab-booking-system/repository"
	"testing"
)

func TestCreateTrip(t *testing.T) {
	repo := repository.NewTripRepository()
	trip := &models.Trip{
		Id:          "trip1",
		RiderID:     "rider1",
		DriverID:    "driver1",
		Source:      models.Location{Latitude: 12.0, Longitude: 77.0},
		Destination: models.Location{Latitude: 12.9, Longitude: 77.5},
		Charges:     500,
	}

	result, err := repo.CreateTrip(trip)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.Id != trip.Id {
		t.Errorf("Expected trip ID %s, got %s", trip.Id, result.Id)
	}
}
