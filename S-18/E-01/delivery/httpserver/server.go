package httpserver

import (
	"E-01/config"
	"E-01/delivery/httpserver/backofficeuserhandler"
	"E-01/delivery/httpserver/userhandler"
	"E-01/service/authorizationservice"
	"E-01/service/authservice"
	"E-01/service/backofficeuserservice"
	"E-01/service/userservice"
	"E-01/validator/uservalidator"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config                config.Config
	userHandler           userhandler.Handler
	backofficeUserHandler backofficeuserhandler.Handler
}

func New(config config.Config, authSvc authservice.Service, userSvc userservice.Service,
	userValidator uservalidator.Validator, backofficeUserSvc backofficeuserservice.Service, authorizationSvc authorizationservice.Service) Server {
	return Server{
		config:                config,
		userHandler:           userhandler.New(authSvc, userSvc, userValidator, config.Auth),
		backofficeUserHandler: backofficeuserhandler.New(config.Auth, authSvc, backofficeUserSvc, authorizationSvc),
	}
}

func (s Server) Serve() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/health-check", s.healthCheck)

	s.userHandler.SetRoutes(e)
	s.backofficeUserHandler.SetRoutes(e)

	// Start Server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.config.HTTPServer.Port)))
}
