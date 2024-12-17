package main

import (
	services "Fitness/Services"
	"fmt"
)

func main() {

	err := services.CreateClass("Mass Gainer", "2024-12-01", "2024-12-20", 10)
	if err != nil {
		fmt.Println("Error creating class:", err)
	}

	err = services.CreateClass("Yoga", "2024-12-05", "2024-12-15", 15)
	if err != nil {
		fmt.Println("Error creating class:", err)
	}

	services.ListClasses()
}
