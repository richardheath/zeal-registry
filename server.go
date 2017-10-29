package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var zeal zealInstance

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

	zeal := zealInstance{}
	err = zeal.initialize(config)
	if err != nil {
		fmt.Println("Failed to initialize zeal registry")
		fmt.Println(err.Error())
		return
	}

	e := echo.New()

	e.Use(middleware.BasicAuth(basicAuthHandler))

	e.GET("/", homeHandler)
	e.GET("/api/v1/:repo/search/:keywords", searchHandler)
	e.GET("/api/v1/:repo/download/:package/:version/:platform", downloadHandler)
	e.GET("/api/v1/:repo/metadata/:package/:version", metadataHandler)
	e.GET("/api/v1/:repo/versions/:package", versionsHandler)
	e.POST("/api/v1/:repo/publish", publishHandler)
	e.Logger.Fatal(e.Start(":9090"))
}

func basicAuthHandler(username, password string, c echo.Context) (bool, error) {
	repo := c.Param("repo")
	if repo == "" {
		return true, nil
	}

	return true, nil
}
