package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func publishHandler(c echo.Context) error {
	repo := c.Param("repo")
	packageName := c.FormValue("name")
	version := c.FormValue("version")
	platform := c.FormValue("platform")

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	err = zeal.storage.PublishPackage(repo, packageName, version, platform, &src)
	if err != nil {
		return err
	}

	err = zeal.data.AddPackagePlatform(packageName, version, platform)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"result": "success",
	})
}
