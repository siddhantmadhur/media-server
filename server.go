package main

import (
	"fmt"
	"log"
	"ocelot/config"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	err := runSqliteInit()
	if err != nil {
		log.Fatal("There was an error in connecting to the sqlite file: " + err.Error())
		os.Exit(1)
	}

	var config config.Config
	err = config.Read()

	if err != nil {
		log.Fatal("[ERROR]: Config could not be read. " + err.Error())
		os.Exit(1)
	}

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		//AllowMethods: []string{"POST", "GET", "UPDATE", "DELETE"},
		AllowHeaders: []string{"Client-Name", "Content-Type", "Authorization"},
	}))

	handler(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
