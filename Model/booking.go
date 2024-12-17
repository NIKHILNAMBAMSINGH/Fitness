package model

import "time"

type Booking struct {
	memberName string
	className  string
	date       time.Time
}
