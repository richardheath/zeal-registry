package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func metadataHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Search\n")
}
