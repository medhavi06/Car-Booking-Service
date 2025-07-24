package services

import (
	"cab-booking-system/models"
	"cab-booking-system/repository/interfaces"
	"errors"
)

type TripService struct {
	tripRepo  interfaces.ITripRepository
	riderRepo interfaces.IRiderInterface
	cabRepo   interfaces.ICabRepository
}

func (s *TripService) CreateTrip(riderId string, source, destination models.Location) (*models.Trip, error) {
	rider, err := s.riderRepo.GetRider(riderId)
	if err != nil {
		return nil, errors.New("invalid rider")
	}
	nearbyAvailableDrivers, err := s.cabRepo.GetAvailableCabsNearby(source, 5)
	if err != nil {
		return nil, errors.New("no driver available")
	}
	cab := nearbyAvailableDrivers[0]
	trip, err := s.tripRepo.CreateTrip(rider.Id, cab.ID, source, destination)
	if err != nil {
		return nil, errors.New("failed to create trip")
	}
	// Set cab's CurrentTripID
	err = s.cabRepo.SetCurrentTripID(cab.ID, trip.ID)
	if err != nil {
		return nil, errors.New("failed to update cab status")
	}
	return trip, nil
}

func (s *TripService) GetTripsByRider(riderID string) ([]*models.Trip, error) {
	allTrips, err := s.tripRepo.GetAllTripsByRider(riderID)
	if err != nil {
		return nil, errors.New("no trip found")
	}
	return allTrips, nil
}

func (s *TripService) EndTrip(tripID string) error {
	trip, err := s.tripRepo.GetTrip(tripID)
	if err != nil {
		return err
	}
	err = s.tripRepo.EndTrip(tripID)
	if err != nil {
		return err
	}
	// Clear cab's CurrentTripID
	return s.cabRepo.SetCurrentTripID(trip.CabID, "")
}

func NewTripService(tripRepo interfaces.ITripRepository, riderRepo interfaces.IRiderInterface, cabRepo interfaces.ICabRepository) *TripService {
	return &TripService{tripRepo: tripRepo, riderRepo: riderRepo, cabRepo: cabRepo}
}
