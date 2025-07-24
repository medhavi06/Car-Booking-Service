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
	createdTrip, err := s.tripRepo.CreateTrip(*rider, *nearbyAvailableDrivers[0], source, destination)
	if err != nil {
		return nil, errors.New("failed to create trip")
	}
	return createdTrip, nil
}

func (s *TripService) GetTripsByRider(riderID string) ([]*models.Trip, error) {
	allTrips, err := s.tripRepo.GetAllTripsByRider(riderID)
	if err != nil {
		return nil, errors.New("no trip found")
	}
	return allTrips, nil
}

func (s *TripService) EndTrip(driver models.Cab) {
	_ = s.tripRepo.EndTrip(driver)
}

func NewTripService(tripRepo interfaces.ITripRepository) *TripService {
	return &TripService{tripRepo: tripRepo}
}
