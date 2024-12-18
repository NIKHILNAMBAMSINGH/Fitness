package services

import (
	model "Fitness/Model"
	"errors"
	"fmt"
	"time"
)

const dateFormatForBooking = "2006-01-02"

var (
	bookings       []model.Booking
	bookingCounter int
)

func GetAllBookings() ([]model.Booking, error) {
	if len(bookings) == 0 {
		return nil, errors.New("no bookings found")
	}
	for _, data := range bookings {
		memberName := data.BookingMemberName
		className := data.BookingClassName
		date := data.BookingDate.Format(dateFormatForBooking)
		fmt.Printf("Member: %s ---> Date: %s ----> Class: %s\n", className, memberName, date)
	}
	return bookings, nil
}

func CreateBooking(name string, className string, bookingDate time.Time) error {

	var selectedClass *model.Class
	for i := range classes {
		class := classes[i]
		if class.Name == className {
			selectedClass = &class
			break
		}
	}
	if selectedClass == nil {
		return fmt.Errorf("No class found for Class: %s", className)
	}
	bookingCounter++
	newBooking := model.Booking{
		Id:                bookingCounter,
		BookingMemberName: name,
		BookingClassName:  className,
		BookingDate:       bookingDate,
	}
	bookings = append(bookings, newBooking)
	fmt.Printf("Booked class for %s on %s for class '%s'.\n", name, newBooking.BookingDate, className)
	return nil

}
