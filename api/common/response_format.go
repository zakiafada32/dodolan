package common

import (
	"net/http"

	"github.com/zakiafada32/retail/business"
)

type ResponseFormat struct {
	StatusCode string      `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func ConstructResponse(status string, data map[string]interface{}) (int, ResponseFormat) {
	var httpStatus int
	var response ResponseFormat

	switch status {
	case business.BadRequest:
		httpStatus = http.StatusBadRequest
		response.StatusCode = business.BadRequest
		response.Message = "bad request"
		response.Data = ""
	case business.InternalServerError:
		httpStatus = http.StatusInternalServerError
		response.StatusCode = business.InternalServerError
		response.Message = "internal server error"
		response.Data = ""
	case business.LoginSuccess:
		httpStatus = http.StatusOK
		response.StatusCode = business.LoginSuccess
		response.Message = "login success"
		response.Data = data
	}

	return httpStatus, response
}
