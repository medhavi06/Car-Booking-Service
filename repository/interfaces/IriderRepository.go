package interfaces

import "cab-booking-system/models"

type IRiderInterface interface {
	AddRider(id, name, phone, email string) error
	GetRider(id string) (*models.Rider, error)
	RateDriver()
	LoadRiderHistory(id string)
}
