package model

import "time"

type Booking struct {
	Id                int
	BookingMemberName string
	BookingClassName  string
	BookingDate       time.Time
}
