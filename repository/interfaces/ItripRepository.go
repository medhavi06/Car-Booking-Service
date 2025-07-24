package interfaces

import "cab-booking-system/models"

type ITripRepository interface {
	CreateTrip(riderID, cabID string, sourceLocation, destinationLocation models.Location) (*models.Trip, error)
	GetAllTripsByRider(riderID string) ([]*models.Trip, error)
	GetTrip(tripID string) (*models.Trip, error)
	EndTrip(tripID string) error
}
