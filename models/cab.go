package models

type Cab struct {
	Id              string
	No              string
	DriverName      string
	PhoneNo         string
	Rating          float32
	CurrentLocation Location
	Availability    bool
	CurrentTrip     *Trip
}
