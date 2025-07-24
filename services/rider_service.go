package services

import (
	"cab-booking-system/models"
	"cab-booking-system/repository/interfaces"
)

type RiderService struct {
	repo interfaces.IRiderInterface
}

func NewRiderService(repo interfaces.IRiderInterface) *RiderService {
	return &RiderService{repo: repo}
}

func (s *RiderService) AddRider(id, name, phone, email string) error {
	return s.repo.AddRider(id, name, phone, email)
}

func (s *RiderService) GetRider(id string) (*models.Rider, error) {
	return s.repo.GetRider(id)
}
