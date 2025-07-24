package interfaces

import "cab-booking-system/models"

type ITripRepository interface {
	CreateTrip(rider models.Rider, cab models.Cab, sourceLocation, destinationLocation models.Location) (*models.Trip, error)
	GetAllTripsByRider(riderID string) ([]*models.Trip, error)
	EndTrip(cab models.Cab) error
}
