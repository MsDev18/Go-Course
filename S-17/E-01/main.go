package main

import (
	"E-01/config"
	"E-01/delivery/httpserver"
	"E-01/repository/mysql"
	"E-01/service/authservice"
	"E-01/service/userservice"
	"E-01/validator/uservalidator"
	"fmt"
	"time"
)

const (
	JwtSignKey                 = "jwt_secret"
	AccessTokenSubject         = "at"
	RefreshTokenSubject        = "rt"
	AccessTokenExpireDuration  = time.Hour * 24
	RefreshTokenExpireDuration = time.Hour * 24 * 7
)

func main() {
	
	cfg := config.Config{
		HTTPServer: config.HTTPServer{Port: 8088},
		Auth: authservice.Config{
			SignKey:               JwtSignKey,
			AccessExpirationTime:  AccessTokenExpireDuration,
			RefreshExpirationTime: RefreshTokenExpireDuration,
			AccessSubject:         AccessTokenSubject,
			RefreshSubject:        RefreshTokenSubject,
		},
		Mysql: mysql.Config{
			Username: "gameapp",
			Password: "gameappt0lk2o20",
			Port:     3308,
			Host:     "localhost",
			DBName:   "gameapp_db",
		},
	}

	// TODO - add command for migration
	// mgr := migrator.New(cfg.Mysql)
	// mgr.Up()

	authSvc, userSvc, userValidator := setupServices(cfg)

	server := httpserver.New(cfg, authSvc, userSvc, userValidator)

	fmt.Println("Start Echo Server ...")
	server.Serve()
}

func setupServices(cfg config.Config) (authservice.Service, userservice.Service, uservalidator.Validator) {
	MysqlRepo := mysql.New(cfg.Mysql)

	uV := uservalidator.New(MysqlRepo)
	authSvc := authservice.New(cfg.Auth)
	userSvc := userservice.New(authSvc, MysqlRepo)


	return authSvc, userSvc ,uV
}
