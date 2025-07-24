package tests

import (
	"cab-booking-system/mocks"
	"cab-booking-system/models"
	"cab-booking-system/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCabService_AddCab_Success(t *testing.T) {
	mockRepo := new(mocks.ICabRepository)
	mockRepo.On("AddCab", "cab1", "DL01AB1234", "John Doe", "9876543210").Return(nil)

	service := services.NewCabService(mockRepo)
	err := service.AddCab("cab1", "DL01AB1234", "John Doe", "9876543210")
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCabService_AddCab_Error(t *testing.T) {
	mockRepo := new(mocks.ICabRepository)
	mockRepo.On("AddCab", "cab1", "DL01AB1234", "John Doe", "9876543210").Return(assert.AnError)

	service := services.NewCabService(mockRepo)
	err := service.AddCab("cab1", "DL01AB1234", "John Doe", "9876543210")
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCabService_UpdateDriverLocation_Success(t *testing.T) {
	mockRepo := new(mocks.ICabRepository)
	loc := &models.Location{Latitude: 12.0, Longitude: 77.0}
	mockRepo.On("UpdateLocation", "cab1", loc).Return(nil)

	service := services.NewCabService(mockRepo)
	err := service.UpdateDriverLocation("cab1", loc)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCabService_UpdateDriverLocation_Error(t *testing.T) {
	mockRepo := new(mocks.ICabRepository)
	loc := &models.Location{Latitude: 12.0, Longitude: 77.0}
	mockRepo.On("UpdateLocation", "cab1", loc).Return(assert.AnError)

	service := services.NewCabService(mockRepo)
	err := service.UpdateDriverLocation("cab1", loc)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}
