package services

import (
	model "Fitness/Model"
	"errors"
	"fmt"
	"time"
)

const dateFormatForBooking = "2006-01-02"

var bookings []model.Booking

func GetAllBookings() {
	for _, data := range bookings {
		memberName := data.BookingMemberName
		className := data.BookingClassName
		date := data.BookingDate.Format(dateFormatForBooking)
		fmt.Printf("Member: %s ---> Date: %s ----> Class: %s\n", className, memberName, date)

	}

}

func CreateBooking(name string, className string, dateString string) error {
	bookedDate, err := time.Parse(dateFormatForBooking, dateString)
	if err != nil {
		return errors.New("Date format is wrong correct format YYYY-MM-dd")
	}
	var selectedClass *model.Class
	for i := range classes {
		class := classes[i]
		if class.Name == className {
			selectedClass = &class
			break
		}
	}
	if selectedClass == nil {
		return errors.New("no class found ")
	}

	newBooking := model.Booking{
		BookingMemberName: name,
		BookingClassName:  className,
		BookingDate:       bookedDate,
	}
	bookings = append(bookings, newBooking)

	fmt.Printf("Booked class for %s on %s for class '%s'.\n", name, bookedDate, className)
	return nil

}
