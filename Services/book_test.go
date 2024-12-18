package services

import (
	model "Fitness/Model"
	"errors"
	"testing"
	"time"
)

func TestCreateBooking(t *testing.T) {
	classes = []model.Class{
		{
			Id:        1,
			Name:      "Yoga",
			StartDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
			Capacity:  10,
		},
		{
			Id:        2,
			Name:      "Gym",
			StartDate: time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2024, 1, 4, 0, 0, 0, 0, time.UTC),
			Capacity:  5,
		},
	}

	type test struct {
		memberName string
		className  string
		date       string
		expected   error
		bookings   []model.Booking
	}

	tests := []test{
		{
			memberName: "Nikhil",
			className:  "Yoga",
			date:       "2024-01-02",
			expected:   nil,
			bookings:   []model.Booking{},
		},
		{
			memberName: "Niraj",
			className:  "Yoga",
			date:       "2024-01-01",
			expected:   errors.New("booking date cannot be in the past"),
			bookings:   []model.Booking{},
		},
		{
			memberName: "Nikhil",
			className:  "Yoga",
			date:       "2024-01-02",
			expected:   errors.New("duplicate booking for the same class and member"),
			bookings: []model.Booking{
				{
					Id:                1,
					BookingMemberName: "Niraj",
					BookingClassName:  "Yoga",
					BookingDate:       time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
				},
			},
		},
	}

	for _, tc := range tests {
		bookings = tc.bookings
		bookingDate, _ := time.Parse("2006-01-02", tc.date)
		err := CreateBooking(tc.memberName, tc.className, bookingDate)
		if err != nil && err.Error() != tc.expected.Error() {
			t.Errorf("expected error: %v, got: %v", tc.expected, err)
		}
	}
}
