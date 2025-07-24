package models

type TripStatus int

// 2. Define constants for the enum values using iota
const (
	TripStatusInProgress TripStatus = iota // 0
	TripStatusFinished                     // 1
)

type Trip struct {
	Id          string
	Rider       Rider
	Driver      Cab
	Source      Location
	Destination Location
	Charges     float32
	Status      TripStatus
}
