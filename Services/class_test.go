package services

import (
	model "Fitness/Model"
	"errors"
	"testing"
	"time"
)

func TestCreateClass(t *testing.T) {
	type test struct {
		name      string
		startDate string
		endDate   string
		capacity  int
		expected  error
		classes   []model.Class
	}

	tests := []test{
		{
			name:      "Zumba",
			startDate: "2024-01-01",
			endDate:   "2024-01-02",
			capacity:  10,
			expected:  nil,
			classes:   []model.Class{},
		},
		{
			name:      "Gym",
			startDate: "2024-01-01",
			endDate:   "2024-01-02",
			capacity:  0,
			expected:  errors.New("Zero capacity is not accepted"),
			classes:   []model.Class{},
		},
		{
			name:      "Zumba",
			startDate: "2024-01-01",
			endDate:   "2024-01-02",
			capacity:  10,
			expected:  errors.New("Class with the same name and date range already exists"),
			classes: []model.Class{
				{
					Id:        1,
					Name:      "Yoga",
					StartDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
					EndDate:   time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
					Capacity:  10,
				},
			},
		},
	}

	for _, tc := range tests {

		classes = tc.classes
		startDate, _ := time.Parse("2006-01-02", tc.startDate)
		endDate, _ := time.Parse("2006-01-02", tc.endDate)

		err := CreateClass(tc.name, startDate, endDate, tc.capacity)
		if err != nil && err.Error() != tc.expected.Error() {
			t.Errorf("expected error: %v, got: %v", tc.expected, err)
		}
	}
}
