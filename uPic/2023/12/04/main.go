package main

import (
	"fmt"
	"net/http"
	"opstools/cloud"

	echo "github.com/labstack/echo/v4"
)

func main() {
	s := cloud.Getissues()
	fmt.Println(s)
	fmt.Println("test")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
