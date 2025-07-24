package tests

import (
	"cab-booking-system/mocks"
	"cab-booking-system/models"
	"cab-booking-system/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRiderService_AddRider_Success(t *testing.T) {
	mockRepo := new(mocks.IRiderInterface)
	mockRepo.On("AddRider", "r1", "Alice", "9999999999", "alice@example.com").Return(nil)

	service := services.NewRiderService(mockRepo)
	err := service.AddRider("r1", "Alice", "9999999999", "alice@example.com")
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRiderService_AddRider_Error(t *testing.T) {
	mockRepo := new(mocks.IRiderInterface)
	mockRepo.On("AddRider", "r1", "Alice", "9999999999", "alice@example.com").Return(assert.AnError)

	service := services.NewRiderService(mockRepo)
	err := service.AddRider("r1", "Alice", "9999999999", "alice@example.com")
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRiderService_GetRider_Success(t *testing.T) {
	mockRepo := new(mocks.IRiderInterface)
	rider := &models.Rider{Id: "r1", Name: "Alice"}
	mockRepo.On("GetRider", "r1").Return(rider, nil)

	service := services.NewRiderService(mockRepo)
	result, err := service.GetRider("r1")
	assert.NoError(t, err)
	assert.Equal(t, rider, result)
	mockRepo.AssertExpectations(t)
}

func TestRiderService_GetRider_Error(t *testing.T) {
	mockRepo := new(mocks.IRiderInterface)
	mockRepo.On("GetRider", "r2").Return(nil, assert.AnError)

	service := services.NewRiderService(mockRepo)
	result, err := service.GetRider("r2")
	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
