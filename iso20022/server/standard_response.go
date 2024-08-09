package server

import (
	"bytes"
)

const (
	OK   ResponseStatus = "Ok"
	Fail ResponseStatus = "Fail"
)

type ResponseStatus string

type StandardResponse struct {
	Status  ResponseStatus `json:"status"`
	Message string         `json:"message"`
	Data    any            `json:"data,omitempty"`
}

func GetFailResponse(msg string, data any) StandardResponse {
	return StandardResponse{
		Status:  Fail,
		Message: msg,
		Data:    data,
	}
}

func GetSuccessResponse(data any) StandardResponse {
	return StandardResponse{
		Status: OK,
		Data:   data,
	}
}

func GetFailResponseFromErrors(errors ...error) StandardResponse {
	buf := bytes.Buffer{}
	for _, err := range errors {
		buf.WriteString(err.Error())
		buf.WriteString("\n")
	}

	return GetFailResponse(buf.String(), nil)
}
