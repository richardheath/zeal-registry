package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func downloadHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Download\n")
}
