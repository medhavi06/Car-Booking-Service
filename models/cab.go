package models

type Cab struct {
	ID              string
	Number          string
	DriverName      string
	PhoneNumber     string
	Rating          float32
	CurrentLocation Location
	Availability    bool
	CurrentTripID   string // empty if not on a trip
}
