package tests

import (
	"cab-booking-system/repository"
	"testing"
)

func TestAddAndGetCab(t *testing.T) {
	repo := repository.NewCabRepository()
	err := repo.AddCab("cab1", "DL01AB1234", "John Doe", "9876543210")
	if err != nil {
		t.Fatalf("AddCab failed: %v", err)
	}
	got, err := repo.("cab1")
	if err != nil {
		t.Fatalf("GetCab failed: %v", err)
	}
	if got.Id != cab.Id {
		t.Errorf("Expected cab ID %v, got %v", cab.Id, got.Id)
	}
}

func TestAddCabDuplicate(t *testing.T) {
	repo := repository.NewCabRepository()
	_ = repo.AddCab("cab1", "DL01AB1234", "John Doe", "9876543210")
	err := repo.AddCab("cab1", "DL01AB1234", "John Doe", "9876543210")
	if err == nil {
		t.Errorf("Expected error for duplicate cab, got nil")
	}
}

func TestGetCabNotFound(t *testing.T) {
	repo := repository.NewCabRepository()
	err := repo.GetCab("notfound")
	if err == nil {
		t.Errorf("Expected error for missing cab, got nil")
	}
}
