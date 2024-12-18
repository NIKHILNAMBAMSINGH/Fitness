package model

import "time"

type Class struct {
	Id        int
	Name      string
	StartDate time.Time
	EndDate   time.Time
	Capacity  int
}
