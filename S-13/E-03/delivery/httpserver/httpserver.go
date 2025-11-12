package httpserver

import (
	"E-03/config"
	"E-03/service/auth"
	"E-03/service/user"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config  config.Config
	authSvc auth.Service
	userSvc user.Service
}

func New(config config.Config, authSvc auth.Service, userSvc user.Service) Server {
	return Server{
		config:  config,
		authSvc: authSvc,
		userSvc: userSvc,
	}
}

func (s Server) Serve() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/health-check", s.healthCheck)

	e.POST("/users/register", s.userRegister)

	// Start Server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.config.HTTPServer.Port)))
}
