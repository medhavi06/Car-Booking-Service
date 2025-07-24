package repository

import (
	"cab-booking-system/models"
	"errors"
	"fmt"
)

type RiderRepository struct {
	riders map[string]*models.Rider
}

func (r *RiderRepository) AddRider(id, name, phone, email string) error {
	if id == "" || name == "" {
		return errors.New("invalid id")
	}
	if riderNode, ok := r.riders[id]; ok {
		fmt.Printf("Rider with ID %d exists! Name: %s\n", id, riderNode.Name)
		return errors.New("rider already exists")
	}
	r.riders[id] = &models.Rider{
		Id:      id,
		Name:    name,
		PhoneNo: phone,
		Email:   email,
		Rating:  0,
	}
	return nil
}

func NewRiderRepository() *RiderRepository {
	return &RiderRepository{riders: make(map[string]*models.Rider)}
}

func (r *RiderRepository) LoadRiderHistory(id string) {

}
func (r *RiderRepository) GetRider(id string) (*models.Rider, error) {
	if rider, exists := r.riders[id]; exists {
		return rider, nil
	}
	return nil, errors.New("rider not found")
}

func (r *RiderRepository) RateDriver() {

}
