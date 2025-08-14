package cmd

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartServer() error {
	e := echo.New()
	// middleware basico de echo
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})
	e.Logger.Info("Server starting on :8080")
	return e.Start(":8080")
}
