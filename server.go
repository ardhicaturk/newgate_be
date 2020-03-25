package main

import (
	"github.com/ardhicaturk/newgate_be/routes"
	"github.com/ardhicaturk/newgate_be/webserver/setupmiddleware"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	setupmiddleware.SetLogMiddleware(e)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:[]string{"*"},
		AllowMethods:[]string{
			echo.GET,
			echo.POST,
			echo.PUT,
			echo.DELETE,
		},
	}))

	apiV1 := e.Group("/api/v1")
	routes.V1(apiV1)

	e.Logger.Fatal(e.Start(":3233"))
}