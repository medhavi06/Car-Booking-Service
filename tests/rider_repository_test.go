package tests

import (
	"cab-booking-system/repository"
	"testing"
)

func TestAddAndGetRider(t *testing.T) {
	repo := repository.NewRiderRepository()
	err := repo.AddRider("r1", "Alice", "9999999999", "alice@example.com")
	if err != nil {
		t.Fatalf("AddRider failed: %v", err)
	}

	rider, err := repo.GetRider("r1")
	if err != nil {
		t.Fatalf("GetRider failed: %v", err)
	}
	if rider.Name != "Alice" {
		t.Errorf("Expected name Alice, got %s", rider.Name)
	}
}
