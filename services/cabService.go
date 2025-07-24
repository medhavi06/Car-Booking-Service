package services

import (
	"cab-booking-system/models"
	"cab-booking-system/repository/interfaces"
)

type CabService struct {
	repo     interfaces.ICabRepository
	tripRepo interfaces.ITripRepository
}

func (s *CabService) UpdateDriverLocation(id string, location *models.Location) {
	_ = s.repo.UpdateLocation(id, location)
}

func NewCabService(repo interfaces.ICabRepository) *CabService {
	return &CabService{repo: repo}
}

func (s *CabService) AddCab(id, no, name, phone string) {
	_ = s.repo.AddCab(id, no, name, phone)
}
