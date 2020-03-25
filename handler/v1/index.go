package v1

import (
	"github.com/labstack/echo"
	"net/http"
)

type ResponseSuccess struct {
	Code 	string `json:"code" xml:"code"`
	Payload interface{} `json:"payload" xml:"payload"`
}

func HandlerIndex(c echo.Context) error {
	return c.JSON(http.StatusOK, ResponseSuccess{
		Code:    "200",
		Payload: struct {
			Name string
			API_Version int
		}{
			Name: "Learn Golang",
			API_Version: 1,
		},
	})
}