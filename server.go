package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

var zeal = zealInstance{}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Second argument must be config file")
		return
	}

	config, err := loadConfigFile(os.Args[1])
	if err != nil {
		fmt.Println("Failed to load config")
		fmt.Println(err.Error())
		return
	}

	err = zeal.initialize(config)
	if err != nil {
		fmt.Println("Failed to initialize zeal registry")
		fmt.Println(err.Error())
		return
	}

	e := echo.New()

	e.HTTPErrorHandler = customHTTPErrorHandler

	e.GET("/", homeHandler)
	// Search
	e.GET("/api/v1/:repo/search/keywords/:keywords", searchKeywordsHandler)
	e.POST("/api/v1/:repo/search/versions", searchPackageVersionsHandler)
	// Package
	e.GET("/api/v1/:repo/:package", packageMetadataHandler)
	e.GET("/api/v1/:repo/:package/:platform/:version/definition", downloadDefinitionHandler)
	e.POST("/api/v1/:repo/:package/:platform/:version/definition", publishDefinitionHandler)
	e.GET("/api/v1/:repo/:package/:platform/:version/package", downdloadPackageHandler)
	e.POST("/api/v1/:repo/:package/:platform/:version/package", publishPackageHandler)
	e.Logger.Fatal(e.Start(":9090"))
}

func basicAuthHandler(username, password string, c echo.Context) (bool, error) {
	repo := c.Param("repo")
	if repo == "" {
		return true, nil
	}

	return true, nil
}

func customHTTPErrorHandler(err error, c echo.Context) {
	fmt.Println(err.Error())
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	errorPage := fmt.Sprintf("%d.html", code)
	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}
	c.Logger().Error(err)
}
