package main

import (
	"E-01/config"
	"E-01/delivery/httpserver"
	"E-01/repository/mysql"
	"E-01/service/auth"
	"E-01/service/user"
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
		Auth: auth.Config{
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

	authSvc, userSvc := setupServices(cfg)

	server := httpserver.New(cfg, authSvc, userSvc)

	fmt.Println("Start Echo Server ...")
	server.Serve()
}

func setupServices(cfg config.Config) (auth.Service, user.Service) {
	MysqlRepo := mysql.New(cfg.Mysql)

	authSvc := auth.New(cfg.Auth)
	userSvc := user.New(authSvc, MysqlRepo)

	return authSvc, userSvc
}
