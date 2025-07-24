package main

import (
	"cab-booking-system/models"
	"cab-booking-system/repository"
	"cab-booking-system/services"
	"fmt"
)

func main() {
	tripRepo := repository.NewTripRepository()
	tripService := services.NewTripService(tripRepo)

	riderRepo := repository.NewRiderRepository()
	riderService := services.NewRiderService(riderRepo)
	riderService.AddRider("r1", "John Doe", "123456789", "john@doe.com")

	cabRepo := repository.NewCabRepository()
	cabService := services.NewCabService(cabRepo)
	cabService.AddCab("d1", "ABC1234WER", "Johny Bravo", "987654321")

	createdTrip, err := tripService.CreateTrip("r1", models.Location{Latitude: 12.9716, Longitude: 77.5946}, models.Location{Latitude: 12.2958, Longitude: 76.6394})
	if err != nil {
		fmt.Println("error creating trip:", err)
	}
	fmt.Println("Trip created successfully: ", createdTrip)

}
