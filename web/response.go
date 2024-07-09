package web

import "net/http"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ErrorResponse(message string) *Response {
	return &Response{
		Status:  http.StatusInternalServerError,
		Message: message,
		Data:    nil,
	}
}

func SuccessResponse(data interface{}) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    data,
	}
}
