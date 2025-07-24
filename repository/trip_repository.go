package repository

import (
	"cab-booking-system/models"
	"errors"
	"sync"

	"github.com/google/uuid"
)

type TripRepository struct {
	trips      map[string]*models.Trip
	riderTrips map[string][]*models.Trip
	cabTrips   map[string][]*models.Trip
	mu         sync.RWMutex
}

func NewTripRepository() *TripRepository {
	return &TripRepository{
		trips:      make(map[string]*models.Trip),
		riderTrips: make(map[string][]*models.Trip),
		cabTrips:   make(map[string][]*models.Trip),
	}
}

func (r *TripRepository) CreateTrip(riderID, cabID string, source, destination models.Location) (*models.Trip, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	tripID := uuid.New().String()
	trip := &models.Trip{
		ID:          tripID,
		RiderID:     riderID,
		CabID:       cabID,
		Source:      source,
		Destination: destination,
		Charges:     0,
		Status:      models.TripStatusInProgress,
	}
	r.trips[tripID] = trip
	r.riderTrips[riderID] = append(r.riderTrips[riderID], trip)
	r.cabTrips[cabID] = append(r.cabTrips[cabID], trip)
	return trip, nil
}

func (r *TripRepository) GetAllTripsByRider(riderID string) ([]*models.Trip, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.riderTrips[riderID], nil
}

func (r *TripRepository) GetTrip(tripID string) (*models.Trip, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	trip, ok := r.trips[tripID]
	if !ok {
		return nil, errors.New("trip not found")
	}
	return trip, nil
}

func (r *TripRepository) EndTrip(tripID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	trip, ok := r.trips[tripID]
	if !ok {
		return errors.New("trip not found")
	}
	trip.Status = models.TripStatusFinished
	return nil
}
