package request

type CreateClassRequest struct {
	Name      string `json:"name" binding:"required"`
	StartDate string `json:"start_date" binding:"required"`
	EndDate   string `json:"end_date" binding:"required"`
	Capacity  int    `json:"capacity" binding:"required"`
}
