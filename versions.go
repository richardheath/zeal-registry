package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func versionsHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Search\n")
}
