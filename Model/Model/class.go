package model

import "time"

type Class struct {
	Name      string
	StartDate time.Time
	EndDate   time.Time
	Capacity  int
}
