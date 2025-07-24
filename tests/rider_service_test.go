package tests

import (
	"cab-booking-system/repository"
	"cab-booking-system/services"
	"testing"
)

func TestRiderService_AddAndGetRider(t *testing.T) {
	repo := repository.NewRiderRepository()
	service := services.NewRiderService(repo)

	err := repo.AddRider("r123", "Alice", "1234567890", "alice@example.com")
	if err != nil {
		t.Fatalf("AddRider failed: %v", err)
	}

	rider, err := repo.GetRider("r123")
	if err != nil {
		t.Fatalf("GetRider failed: %v", err)
	}

	if rider.Id != "r123" {
		t.Errorf("Expected rider ID r123, got %s", rider.Id)
	}
	if rider.Name != "Alice" {
		t.Errorf("Expected name Alice, got %s", rider.Name)
	}
}
