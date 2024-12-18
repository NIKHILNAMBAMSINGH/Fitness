package services

import (
	model "Fitness/Model"
	"errors"
	"fmt"
	"time"
)

const dateFormat = "2006-01-02"

var classes []model.Class

func ListClasses() ([]model.Class, error) {
	if len(classes) == 0 {
		return nil, errors.New("no classes available")
	}

	fmt.Println("All the Classes Available:")
	for _, data := range classes {
		name := data.Name
		capacity := data.Capacity
		startDate := data.StartDate.Format(dateFormat)
		endDate := data.EndDate.Format(dateFormat)
		fmt.Printf("Name: %s ---> Start Date: %s ---> End Date: %s ---> Capacity: %d\n", name, startDate, endDate, capacity)
	}

	return classes, nil
}

func CreateClass(name string, startDate time.Time, endDate time.Time, capacity int) error {

	for _, class := range classes {
		if class.Name == name {
			return errors.New("Class with the same name exists")
		}
	}
	if capacity < 0 {
		return errors.New("Zero capacity is not accepted")
	}

	classNew := model.Class{
		Id:        len(classes) + 1,
		Name:      name,
		StartDate: startDate,
		EndDate:   endDate,
		Capacity:  capacity,
	}
	classes = append(classes, classNew)
	fmt.Printf("Created Class %s\n", name)
	return nil

}
