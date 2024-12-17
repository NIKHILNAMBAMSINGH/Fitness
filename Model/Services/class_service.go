package services

import (
	model "Fitness/Model"
	"errors"
	"fmt"
	"time"
)

const dateFormat = "2006-01-02"

var classes []model.Class

func ListClasses() {
	fmt.Println("All the Classes Available")
	for _, data := range classes {
		name := data.Name
		capacity := data.Capacity
		startDate := data.StartDate.Format(dateFormat)
		endDate := data.EndDate.Format(dateFormat)
		fmt.Printf("Name : %s ---> Start Date : %s ---> End Date : %s ---> Capacity : %d\n", name, startDate, endDate, capacity)
	}
}
func CreateClass(name string, startDateString string, endDateString string, capacity int) error {

	startDate, err := time.Parse(dateFormat, startDateString)
	if err != nil {
		return errors.New("Invalid start date format use correct format YYYY-MM-DD")
	}
	endDate, err := time.Parse(dateFormat, endDateString)
	if err != nil {
		return errors.New("Invalid end date format use the correct format YYYY")
	}

	if capacity < 0 {
		return errors.New("Zero capacity is not accepted")
	}

	classNew := model.Class{
		Name:      name,
		StartDate: startDate,
		EndDate:   endDate,
		Capacity:  capacity,
	}
	classes = append(classes, classNew)
	fmt.Printf("Created Class %s\n", name)
	return nil

}
