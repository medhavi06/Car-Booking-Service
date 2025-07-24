package repository

import (
	"cab-booking-system/models"
	"github.com/google/uuid"
)

var uniqueIdGenerator uuid.UUID

type TripRepository struct {
	trips       map[string][]*models.Trip
	riderTrips  map[string][]*models.Trip
	driverTrips map[string][]*models.Trip
}

func (r *TripRepository) GetAllTripsByRider(riderID string) ([]*models.Trip, error) {
	return r.riderTrips[riderID], nil
}

func NewTripRepository() *TripRepository {
	uniqueIdGenerator = uuid.New()
	return &TripRepository{trips: make(map[string][]*models.Trip)}
}

func (r *TripRepository) CreateTrip(rider models.Rider, cab models.Cab, sourceLocation, destinationLocation models.Location) (*models.Trip, error) {
	trip := &models.Trip{
		Id:          uniqueIdGenerator.String(),
		Rider:       rider,
		Driver:      cab,
		Source:      sourceLocation,
		Destination: destinationLocation,
		Charges:     0,
		Status:      models.TripStatusInProgress,
	}
	r.riderTrips[rider.Id] = append(r.trips[rider.Id], trip)
	return trip, nil
}

func (r *TripRepository) EndTrip(cab models.Cab) error {
	cab.CurrentTrip.Status = models.TripStatusFinished
	cab.CurrentTrip = nil
	return nil
}
