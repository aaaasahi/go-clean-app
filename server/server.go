package server

import (
	"log"

	"go-clean-app/config"

	"go-clean-app/server/route"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)
type Server struct {
	Echo        *echo.Echo
	Router      *route.Router
}

func NewServer(router *route.Router) *Server {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return &Server{
		Echo:  e,
		Router: router,
	}
}

func (s *Server) Run(conf *config.Conf) {
	s.Router.InitRoute(s.Echo)
	
	log.Println("Server starting on port 8080")

	if err := s.Echo.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}