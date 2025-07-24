package interfaces

import "cab-booking-system/models"

type ICabRepository interface {
	GetAvailableCabsNearby(loc models.Location, radiusKm float64) ([]*models.Cab, error)
	UpdateLocation(cabID string, location *models.Location) error
	AddCab(id, number, driverName, phoneNumber string) error
	GetCab(id string) (*models.Cab, error)
	SetCurrentTripID(cabID, tripID string) error
}
