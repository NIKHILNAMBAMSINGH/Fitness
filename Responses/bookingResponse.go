package response

import "time"

type BookingResponse struct {
	ID         int       `json:"id"`
	MemberName string    `json:"member_name"`
	ClassName  string    `json:"class_name"`
	Date       time.Time `json:"date"`
}
