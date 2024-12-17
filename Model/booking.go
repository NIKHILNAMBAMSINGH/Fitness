package model

import "time"

type Booking struct {
	BookingMemberName string
	BookingClassName  string
	BookingDate       time.Time
}
