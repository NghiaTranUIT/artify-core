package model

import "net/http"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewSuccessReponse(payload interface{}) Response {
	response := Response{
		Code:    http.StatusOK,
		Data:    payload,
		Message: "Success",
	}
	return response
}

func NewErrorResponse(err error) Response {
	response := Response{
		Code:    http.StatusBadRequest,
		Data:    err,
		Message: "Error",
	}
	return response
}
