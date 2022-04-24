package response

import "strings"

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

func BuildResponse(status bool, message string, data interface{}) Response {
	response := Response{
		Status:  status,
		Message: message,
		Error:   nil,
		Data:    data,
	}
	return response
}

func BuildErrorResponse(message string, errors string, data interface{}) Response {
	splittedErr := strings.Split(errors, "\n")
	response := Response{
		Status:  false,
		Message: message,
		Error:   splittedErr,
		Data:    data,
	}
	return response
}
