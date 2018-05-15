package utils

import (
	"github.com/NghiaTranUIT/artify-core/model"
	"net/http"
)

func ResponseSuccess(data interface{}) model.Response {
	response := model.Response{
		Code:    http.StatusOK,
		Data:    data,
		Message: "Success",
	}
	return response
}

func ResponseError(err error) model.Response {
	response := model.Response{
		Code:    http.StatusBadRequest,
		Data:    err,
		Message: "Error",
	}
	return response
}
