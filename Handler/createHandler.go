package handler

import (
	model "Fitness/Model"
	request "Fitness/Request"
	response "Fitness/Responses"
	service "Fitness/Services"
	"encoding/json"
	"net/http"
	"time"
)

const dateFormat = "2006-01-02"

func CreateClassHandler(w http.ResponseWriter, r *http.Request) {
	var request request.CreateClassRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response := model.ApiResponse{
			Success: false,
			Message: "Invalid input",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	startDate, err := time.Parse(dateFormat, request.StartDate)
	if err != nil {
		response := model.ApiResponse{
			Success: false,
			Message: "Invalid start date format. Use YYYY-MM-DD.",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	endDate, err := time.Parse(dateFormat, request.EndDate)
	if err != nil {
		response := model.ApiResponse{
			Success: false,
			Message: "Invalid end date format. Use YYYY-MM-DD.",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	if err := service.CreateClass(request.Name, startDate, endDate, request.Capacity); err != nil {
		response := model.ApiResponse{
			Success: false,
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := model.ApiResponse{
		Success: true,
		Message: "Class created successfully",
		Data: response.ClassResponse{
			Name:      request.Name,
			StartDate: startDate,
			EndDate:   endDate,
			Capacity:  request.Capacity,
		},
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func GetClassHandler(w http.ResponseWriter, r *http.Request) {
	classes, err := service.ListClasses()
	if err != nil {
		response := model.ApiResponse{
			Success: false,
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(response)
		return
	}
	var classResponses []response.ClassResponse
	for _, class := range classes {
		classResponses = append(classResponses, response.ClassResponse{
			Name:      class.Name,
			StartDate: class.StartDate,
			EndDate:   class.EndDate,
			Capacity:  class.Capacity,
		})
	}
	resp := model.ApiResponse{
		Success: true,
		Message: "Classes retrieved successfully",
		Data:    classResponses,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
