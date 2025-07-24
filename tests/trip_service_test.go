package tests

import (
	"cab-booking-system/mocks"
	"cab-booking-system/models"
	"cab-booking-system/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTripService_CreateTrip_Success(t *testing.T) {
	mockTripRepo := new(mocks.ITripRepository)
	mockRiderRepo := new(mocks.IRiderInterface)
	mockCabRepo := new(mocks.ICabRepository)

	rider := &models.Rider{Id: "r1", Name: "Alice"}
	cab := &models.Cab{ID: "cab1"}
	trip := &models.Trip{ID: "trip1", RiderID: "r1", CabID: "cab1"}

	mockRiderRepo.On("GetRider", "r1").Return(rider, nil)
	mockCabRepo.On("GetAvailableCabsNearby", mock.Anything, mock.Anything).Return([]*models.Cab{cab}, nil)
	mockTripRepo.On("CreateTrip", "r1", "cab1", mock.Anything, mock.Anything).Return(trip, nil)
	mockCabRepo.On("SetCurrentTripID", "cab1", "trip1").Return(nil)

	service := services.NewTripService(mockTripRepo, mockRiderRepo, mockCabRepo)
	result, err := service.CreateTrip("r1", models.Location{}, models.Location{})
	assert.NoError(t, err)
	assert.Equal(t, trip, result)
	mockTripRepo.AssertExpectations(t)
	mockRiderRepo.AssertExpectations(t)
	mockCabRepo.AssertExpectations(t)
}

func TestTripService_CreateTrip_Error(t *testing.T) {
	mockTripRepo := new(mocks.ITripRepository)
	mockRiderRepo := new(mocks.IRiderInterface)
	mockCabRepo := new(mocks.ICabRepository)

	mockRiderRepo.On("GetRider", "r1").Return(nil, assert.AnError)

	service := services.NewTripService(mockTripRepo, mockRiderRepo, mockCabRepo)
	result, err := service.CreateTrip("r1", models.Location{}, models.Location{})
	assert.Error(t, err)
	assert.Nil(t, result)
	mockRiderRepo.AssertExpectations(t)
}

func TestTripService_GetTripsByRider_Success(t *testing.T) {
	mockTripRepo := new(mocks.ITripRepository)
	mockRiderRepo := new(mocks.IRiderInterface)
	mockCabRepo := new(mocks.ICabRepository)

	trips := []*models.Trip{{ID: "trip1"}}
	mockTripRepo.On("GetAllTripsByRider", "r1").Return(trips, nil)

	service := services.NewTripService(mockTripRepo, mockRiderRepo, mockCabRepo)
	result, err := service.GetTripsByRider("r1")
	assert.NoError(t, err)
	assert.Equal(t, trips, result)
	mockTripRepo.AssertExpectations(t)
}

func TestTripService_GetTripsByRider_Error(t *testing.T) {
	mockTripRepo := new(mocks.ITripRepository)
	mockRiderRepo := new(mocks.IRiderInterface)
	mockCabRepo := new(mocks.ICabRepository)

	mockTripRepo.On("GetAllTripsByRider", "r2").Return(nil, assert.AnError)

	service := services.NewTripService(mockTripRepo, mockRiderRepo, mockCabRepo)
	result, err := service.GetTripsByRider("r2")
	assert.Error(t, err)
	assert.Nil(t, result)
	mockTripRepo.AssertExpectations(t)
}

func TestTripService_EndTrip(t *testing.T) {
	mockTripRepo := new(mocks.ITripRepository)
	mockRiderRepo := new(mocks.IRiderInterface)
	mockCabRepo := new(mocks.ICabRepository)

	trip := &models.Trip{ID: "trip1", CabID: "cab1"}
	mockTripRepo.On("GetTrip", "trip1").Return(trip, nil)
	mockTripRepo.On("EndTrip", "trip1").Return(nil)
	mockCabRepo.On("SetCurrentTripID", "cab1", "").Return(nil)

	service := services.NewTripService(mockTripRepo, mockRiderRepo, mockCabRepo)
	err := service.EndTrip("trip1")
	assert.NoError(t, err)
	mockTripRepo.AssertExpectations(t)
	mockCabRepo.AssertExpectations(t)
}
