package tests

import (
	"cab-booking-system/repository"
	"cab-booking-system/services"
	"testing"
)

func TestCabService_AddAndGetCab(t *testing.T) {
	repo := repository.NewCabRepository()
	service := services.NewCarService(repo)

	cab, err := repo.AddCab("cab001", "KA05MQ1234", "Bob", "9876543210")
	if err != nil {
		t.Fatalf("AddCab failed: %v", err)
	}

	fetchedCab, err := repo.GetCab("cab001")
	if err != nil {
		t.Fatalf("GetCab failed: %v", err)
	}

	if fetchedCab.No != "KA05MQ1234" {
		t.Errorf("Expected cab number KA05MQ1234, got %s", fetchedCab.No)
	}
	if fetchedCab.DriverName != "Bob" {
		t.Errorf("Expected driver Bob, got %s", fetchedCab.DriverName)
	}
}
