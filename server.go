package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {

	err := runSqliteInit()

	if err != nil {
		log.Fatal("There was an error in connecting to the sqlite file: " + err.Error())
		os.Exit(1)
	}

	e := echo.New()

	e.GET("/health-check", func(c echo.Context) error {
		return c.String(http.StatusOK, "Health OK!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
