package httpserver

import (
	"E-01/config"
	"E-01/delivery/httpserver/backofficeuserhandler"
	"E-01/delivery/httpserver/matchinghandler"
	"E-01/delivery/httpserver/userhandler"
	"E-01/service/authorizationservice"
	"E-01/service/authservice"
	"E-01/service/backofficeuserservice"
	"E-01/service/matchingservice"
	"E-01/service/userservice"
	"E-01/validator/matchingvalidator"
	"E-01/validator/uservalidator"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config                config.Config
	userHandler           userhandler.Handler
	backofficeUserHandler backofficeuserhandler.Handler
	matchingHandler       matchinghandler.Handler
	Router *echo.Echo
}

func New(config config.Config, authSvc authservice.Service, userSvc userservice.Service,
	userValidator uservalidator.Validator, backofficeUserSvc backofficeuserservice.Service,
	authorizationSvc authorizationservice.Service, matchingSvc matchingservice.Service, matchingValidator matchingvalidator.Validator) Server {
	return Server{
		Router: echo.New(),
		config:                config,
		userHandler:           userhandler.New(authSvc, userSvc, userValidator, config.Auth),
		backofficeUserHandler: backofficeuserhandler.New(config.Auth, authSvc, backofficeUserSvc, authorizationSvc),
		matchingHandler:       matchinghandler.New(config.Auth, authSvc, matchingSvc, matchingValidator),
	}
}

func (s Server) Serve() {

	// Middleware
	s.Router.Use(middleware.Logger())
	s.Router.Use(middleware.Recover())

	// Routes
	s.Router.GET("/health-check", s.healthCheck)

	s.userHandler.SetRoutes(s.Router)
	s.backofficeUserHandler.SetRoutes(s.Router)
	s.matchingHandler.SetRoutes(s.Router)

	// Start Server

	address := s.config.HTTPServer.Port
	fmt.Printf("Start Echo Server On Port %d \n", address)
	if err := s.Router.Start(fmt.Sprintf(":%d", address)) ; err != nil {
		fmt.Println("router start error : ", err)
	}
}
