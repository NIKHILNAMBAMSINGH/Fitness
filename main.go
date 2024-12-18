package main

import (
	handler "Fitness/Handler"
	"fmt"
	"net/http"
)

func handleClass(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handler.CreateClassHandler(w, r)
	case http.MethodGet:
		handler.GetClassHandler(w, r)
	default:
		http.Error(w, "Only get and post method are allowed", http.StatusMethodNotAllowed)
	}
}

func handleBooking(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handler.CreateBookingHandler(w, r)
	case http.MethodGet:
		handler.GetAllBookingsHandler(w, r)
	default:
		http.Error(w, "Only get and post method are allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", handleClass)
	fmt.Println("Runnning on localhost:8080")
	http.ListenAndServe(":8080", nil)
}
