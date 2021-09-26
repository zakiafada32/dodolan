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
	// success response
	case business.Success:
		httpStatus = http.StatusOK
		response.StatusCode = business.Success
		response.Message = "success"
		response.Data = data
	case business.SucessCreated:
		httpStatus = http.StatusCreated
		response.StatusCode = business.SucessCreated
		response.Message = "success created"
		response.Data = data
	// error response
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
	case business.NotFound:
		httpStatus = http.StatusNotFound
		response.StatusCode = business.NotFound
		response.Message = "data not found"
		response.Data = ""
	case business.Unauthorized:
		httpStatus = http.StatusUnauthorized
		response.StatusCode = business.Unauthorized
		response.Message = "unauthorized"
		response.Data = ""
	}
	return httpStatus, response
}
