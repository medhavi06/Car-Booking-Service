package tests

import (
	"cab-booking-system/models"
	"cab-booking-system/repository"
	"cab-booking-system/services"
	"testing"
)

func TestCreateTripSuccess(t *testing.T) {
	repo := repository.NewTripRepository()
	service := services.NewTripService(repo)

	trip := &models.Trip{
		Id:          "trip2",
		RiderID:     "r1",
		DriverID:    "d1",
		Source:      models.Location{Latitude: 11.0, Longitude: 78.0},
		Destination: models.Location{Latitude: 11.5, Longitude: 78.5},
		Charges:     600,
	}

	createdTrip, err := service.CreateTrip(trip)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if createdTrip == nil || createdTrip.Id != "trip2" {
		t.Errorf("Trip not created as expected")
	}
}

func TestCreateTripNilInput(t *testing.T) {
	repo := repository.NewTripRepository()
	service := services.NewTripService(repo)

	_, err := service.CreateTrip(nil)
	if err == nil {
		t.Errorf("Expected error for nil input, got nil")
	}
}
