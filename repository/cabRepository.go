package repository

import (
	"cab-booking-system/models"
	"errors"
	"math"
)

type CabRepository struct {
	cabs map[string]*models.Cab
}

func (r *CabRepository) GetAvailableCabsNearby(loc models.Location, radiusKm float64) ([]*models.Cab, error) {
	cabs := make([]*models.Cab, 0)
	for _, cab := range r.cabs {
		if cab.Availability && calculateDistance(cab.CurrentLocation, loc) <= radiusKm && cab.CurrentTrip == nil {
			cabs = append(cabs, cab)
		}
	}

	return cabs, nil
}

func NewCabRepository() *CabRepository {
	return &CabRepository{cabs: make(map[string]*models.Cab)}
}

func (r *CabRepository) AddCab(id, no, driverName, phoneNo string) error {
	if id == "" || no == "" || driverName == "" || phoneNo == "" {
		return errors.New("invalid cab details")
	}
	cab := &models.Cab{
		Id:         id,
		No:         no,
		DriverName: driverName,
		PhoneNo:    phoneNo,
	}
	r.cabs[id] = cab
	return nil
}

func (r *CabRepository) UpdateLocation(riderID string, location *models.Location) error {
	if _, ok := r.cabs[riderID]; !ok {
		return errors.New("no driver found")
	}
	r.cabs[riderID].CurrentLocation = *location
	return nil
}

func calculateDistance(driverLocation, riderLocation models.Location) float64 {
	return math.Sqrt(math.Pow(driverLocation.Latitude-riderLocation.Latitude, 2) + math.Pow(driverLocation.Longitude-riderLocation.Longitude, 2))
}
