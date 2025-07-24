package interfaces

import "cab-booking-system/models"

type ICabRepository interface {
	GetAvailableCabsNearby(loc models.Location, radiusKm float64) ([]*models.Cab, error)
	UpdateLocation(riderID string, location *models.Location) error
	AddCab(id, no, driverName, phoneNo string) error
}
