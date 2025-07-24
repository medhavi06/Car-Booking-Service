package repository

import (
	"cab-booking-system/models"
	"errors"
	"math"
	"sync"
)

type CabRepository struct {
	cabs map[string]*models.Cab
	mu   sync.RWMutex
}

func (r *CabRepository) GetAvailableCabsNearby(loc models.Location, radiusKm float64) ([]*models.Cab, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	cabs := make([]*models.Cab, 0)
	for _, cab := range r.cabs {
		if cab.Availability && cab.CurrentTripID == "" && calculateDistance(cab.CurrentLocation, loc) <= radiusKm {
			cabs = append(cabs, cab)
		}
	}
	return cabs, nil
}

func NewCabRepository() *CabRepository {
	return &CabRepository{cabs: make(map[string]*models.Cab, 0)}
}

func (r *CabRepository) AddCab(id, number, driverName, phoneNumber string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if id == "" || number == "" || driverName == "" || phoneNumber == "" {
		return errors.New("invalid cab details")
	}
	if _, ok := r.cabs[id]; ok {
		return errors.New("cab already exists")
	}
	cab := &models.Cab{
		ID:          id,
		Number:      number,
		DriverName:  driverName,
		PhoneNumber: phoneNumber,
	}
	r.cabs[id] = cab
	return nil
}

func (r *CabRepository) GetCab(id string) (*models.Cab, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if id == "" {
		return nil, errors.New("invalid cab details")
	}
	cab, ok := r.cabs[id]
	if !ok {
		return nil, errors.New("cab not found")
	}
	return cab, nil
}

func (r *CabRepository) UpdateLocation(cabID string, location *models.Location) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	cab, ok := r.cabs[cabID]
	if !ok {
		return errors.New("no cab found")
	}
	cab.CurrentLocation = *location
	return nil
}

func (r *CabRepository) SetCurrentTripID(cabID, tripID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	cab, ok := r.cabs[cabID]
	if !ok {
		return errors.New("cab not found")
	}
	cab.CurrentTripID = tripID
	return nil
}

func calculateDistance(driverLocation, riderLocation models.Location) float64 {
	return math.Sqrt(math.Pow(driverLocation.Latitude-riderLocation.Latitude, 2) + math.Pow(driverLocation.Longitude-riderLocation.Longitude, 2))
}
