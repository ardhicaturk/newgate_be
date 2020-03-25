package routes

import (
	v1 "github.com/ardhicaturk/learn_golang/handler/v1"
	"github.com/labstack/echo"
)

func V1(e *echo.Group)  {
	e.GET("/", v1.HandlerIndex)
}