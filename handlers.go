package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func searchKeywordsHandler(c echo.Context) error {
	return successResult(c)
}

func searchPackageVersionsHandler(c echo.Context) error {
	return successResult(c)
}

func packageMetadataHandler(c echo.Context) error {
	return successResult(c)
}

func downloadDefinitionHandler(c echo.Context) error {
	return successResult(c)
}

func publishDefinitionHandler(c echo.Context) error {

	return successResult(c)
}

func downdloadPackageHandler(c echo.Context) error {
	return successResult(c)
}

func publishPackageHandler(c echo.Context) error {
	repo := c.Param("repo")
	err := authenticateWithWrite(repo, c)
	if err != nil {
		return err
	}

	packageName := c.Param("package")
	version := c.Param("version")
	platform := c.Param("platform")

	file, err := c.FormFile("package")
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

	return successResult(c)
}

func successResult(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"result": "success",
	})
}
