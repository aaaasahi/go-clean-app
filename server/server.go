package server

import (
	"log"

	"go-clean-app/config"
	"go-clean-app/server/route"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run(conf *config.Conf) {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	route.InitRoute(e)
	
	log.Println("Server starting on port 8080")

	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}