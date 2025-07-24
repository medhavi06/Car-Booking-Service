package tests

import (
	"cab-booking-system/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRiderRepository_AddRider(t *testing.T) {
	repo := repository.NewRiderRepository()
	err := repo.AddRider("r1", "Alice", "9999999999", "alice@example.com")
	assert.NoError(t, err)

	err = repo.AddRider("r1", "Alice", "9999999999", "alice@example.com")
	assert.Error(t, err)
}

func TestRiderRepository_GetRider(t *testing.T) {
	repo := repository.NewRiderRepository()
	_, err := repo.GetRider("notfound")
	assert.Error(t, err)

	_ = repo.AddRider("r1", "Alice", "9999999999", "alice@example.com")
	rider, err := repo.GetRider("r1")
	assert.NoError(t, err)
	assert.Equal(t, "Alice", rider.Name)
}
