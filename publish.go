package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func publishHandler(c echo.Context) error {
	repo := c.Param("repo")
	err := authenticateWithWrite(repo, c)
	if err != nil {
		return err
	}

	fmt.Println("authorized")
	packageName := c.FormValue("name")
	version := c.FormValue("version")
	platform := c.FormValue("platform")

	fmt.Println("form")
	file, err := c.FormFile("package")
	if err != nil {
		fmt.Println("err: " + err.Error())
		return err
	}

	fmt.Println("here")
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

func authenticateWithWrite(repo string, c echo.Context) error {
	repoSettings, err := zeal.auth.GetRepoSettings(repo)
	if err != nil {
		return err
	}

	if repoSettings.AllowAnonymousWrite {
		return nil
	}

	username, apikey, _ := c.Request().BasicAuth()
	success, err := zeal.auth.Authenticate(username, apikey)
	if err != nil {
		return err
	}

	if !success {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error": "Not authorized",
		})
	}

	return nil
}
