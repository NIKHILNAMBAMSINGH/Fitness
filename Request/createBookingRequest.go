package request

type CreateBookingRequest struct {
	MemberName string `json:"member_name" binding:"required"`
	ClassName  string `json:"class_name" binding:"required"`
	Date       string `json:"date" binding:"required"`
}
