package handler

import (
	"encoding/json"
	"net/http"
	"time"

	model "Fitness/Model"
	request "Fitness/Request"
	response "Fitness/Responses"
	service "Fitness/Services"
)

const dateFormatForBooking = "2006-01-02"

func CreateBookingHandler(w http.ResponseWriter, r *http.Request) {
	var req request.CreateBookingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		resp := model.ApiResponse{
			Success: false,
			Message: "Invalid input",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	}

	bookingDate, err := time.Parse(dateFormatForBooking, req.Date)
	if err != nil {
		resp := model.ApiResponse{
			Success: false,
			Message: "Invalid date format. Use YYYY-MM-DD.",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	}

	if err := service.CreateBooking(req.MemberName, req.ClassName, bookingDate); err != nil {
		resp := model.ApiResponse{
			Success: false,
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(resp)
		return
	}

	bookings, err := service.GetAllBookingss()
	if err != nil {
		resp := model.ApiResponse{
			Success: false,
			Message: "Failed to retrieve bookings",
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(resp)
		return
	}

	newBooking := bookings[len(bookings)-1]
	resp := model.ApiResponse{
		Success: true,
		Message: "Booking created successfully",
		Data: response.BookingResponse{
			ID:         newBooking.Id,
			MemberName: newBooking.BookingMemberName,
			ClassName:  newBooking.BookingClassName,
			Date:       newBooking.BookingDate,
		},
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func GetAllBookingsHandler(w http.ResponseWriter, r *http.Request) {
	bookings, err := service.GetAllBookingss()
	if err != nil {
		response := model.ApiResponse{
			Success: false,
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	var bookingResponses []response.BookingResponse
	for _, booking := range bookings {
		bookingResponses = append(bookingResponses, response.BookingResponse{
			MemberName: booking.BookingMemberName,
			ClassName:  booking.BookingClassName,
			Date:       booking.BookingDate,
		})
	}

	response := model.ApiResponse{
		Success: true,
		Message: "Bookings retrieved successfully",
		Data:    bookingResponses,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
