package model

import (
	"net/http"

	"github.com/labstack/echo"
)

type response struct {
	code    int
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	err     error
}

func BuildResponse(c echo.Context, data interface{}, err error) error {
	response := new(response)

	response.err = err
	response.Data = data

	response.code = http.StatusOK
	response.Message = "OK"

	if response.err != nil {
		response.code = http.StatusInternalServerError
		response.Message = "FAIL"
		response.Data = nil
	}

	return c.JSON(
		response.code,
		response,
	)
}
