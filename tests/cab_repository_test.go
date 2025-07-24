package tests

import (
	"cab-booking-system/models"
	"cab-booking-system/repository"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCabRepository_AddCab(t *testing.T) {
	repo := repository.NewCabRepository()
	err := repo.AddCab("cab1", "DL01AB1234", "John Doe", "9876543210")
	assert.NoError(t, err)

	err = repo.AddCab("cab1", "DL01AB1234", "John Doe", "9876543210")
	assert.Error(t, err)
}

func TestCabRepository_GetCab(t *testing.T) {
	repo := repository.NewCabRepository()
	_, err := repo.GetCab("notfound")
	assert.Error(t, err)

	_ = repo.AddCab("cab1", "DL01AB1234", "John Doe", "9876543210")
	cab, err := repo.GetCab("cab1")
	assert.NoError(t, err)
	assert.Equal(t, "cab1", cab.ID)
}

func TestCabRepository_UpdateLocation(t *testing.T) {
	repo := repository.NewCabRepository()
	_ = repo.AddCab("cab1", "DL01AB1234", "John Doe", "9876543210")
	loc := &models.Location{Latitude: 12.0, Longitude: 77.0}
	err := repo.UpdateLocation("cab1", loc)
	assert.NoError(t, err)

	cab, _ := repo.GetCab("cab1")
	assert.Equal(t, *loc, cab.CurrentLocation)

	err = repo.UpdateLocation("notfound", loc)
	assert.Error(t, err)
}

func TestCabRepository_GetAvailableCabsNearby(t *testing.T) {
	repo := repository.NewCabRepository()
	_ = repo.AddCab("cab1", "DL01AB1234", "John Doe", "9876543210")
	_ = repo.AddCab("cab2", "DL01AB5678", "Jane Doe", "9876543211")
	loc1 := &models.Location{Latitude: 12.0, Longitude: 77.0}
	loc2 := &models.Location{Latitude: 12.1, Longitude: 77.1}
	_ = repo.UpdateLocation("cab1", loc1)
	_ = repo.UpdateLocation("cab2", loc2)
	cab1, _ := repo.GetCab("cab1")
	cab2, _ := repo.GetCab("cab2")
	cab1.Availability = true
	cab2.Availability = true
	cabs, err := repo.GetAvailableCabsNearby(*loc1, 5)
	assert.NoError(t, err)
	assert.NotEmpty(t, cabs)
}

func TestCabRepository_ConcurrentAccess(t *testing.T) {
	repo := repository.NewCabRepository()
	wg := sync.WaitGroup{}
	cabCount := 10

	// Add cabs concurrently
	for i := 0; i < cabCount; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			id := "cab" + string(rune(i))
			repo.AddCab(id, "DL01AB1234", "Driver", "9876543210")
		}(i)
	}

	wg.Wait()

	// Update locations concurrently
	for i := 0; i < cabCount; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			id := "cab" + string(rune(i))
			loc := &models.Location{Latitude: float64(i), Longitude: float64(i)}
			repo.UpdateLocation(id, loc)
		}(i)
	}
	wg.Wait()

	// Read cabs concurrently
	for i := 0; i < cabCount; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			id := "cab" + string(rune(i))
			cab, err := repo.GetCab(id)
			if err == nil {
				assert.Equal(t, id, cab.ID)
			}
		}(i)
	}
	wg.Wait()
}
