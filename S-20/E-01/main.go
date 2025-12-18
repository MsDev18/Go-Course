package main

import (
	"E-01/adapter/redis"
	"E-01/config"
	"E-01/delivery/httpserver"
	"E-01/repository/migrator"
	"E-01/repository/mysql"
	"E-01/repository/mysql/mysqlaccesscontrol"
	"E-01/repository/mysql/mysqluser"
	"E-01/repository/redis/redismatching"
	"E-01/scheduler"
	"E-01/service/authorizationservice"
	"E-01/service/authservice"
	"E-01/service/backofficeuserservice"
	"E-01/service/matchingservice"
	"E-01/service/userservice"
	"E-01/validator/matchingvalidator"
	"E-01/validator/uservalidator"
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	// TODO - read config path from command-line
	cfg := config.Load("config.yml")
	fmt.Printf("cfg : %+v\n", cfg)
	// TODO - merge cfg to cfg2

	// TODO - add command for migration
	mgr := migrator.New(cfg.Mysql)
	mgr.Up()
	// TODO - add stuct and add these returned items as struct field
	authSvc, userSvc, userValidator, backofficeSvc, authorizationSvc, matchingSvc, matchingV := setupServices(cfg)

	var httpServer *echo.Echo
	go func() {
		server := httpserver.New(cfg, authSvc, userSvc, userValidator, backofficeSvc, authorizationSvc, matchingSvc, matchingV)
		httpServer = server.Serve()
	}()

	done := make(chan bool)

	go func() {
		sch := scheduler.New()
		sch.Start(done)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	if err:=httpServer.Shutdown(context.Background()); err != nil {
		fmt.Println("HTTP Server Shotdown Error : ", err)
	}
	fmt.Println("recived interrupt signal, shotting down gracefully... ")
	done <- true
	time.Sleep(time.Second * 5)

}

func setupServices(cfg config.Config) (authservice.Service, userservice.Service, uservalidator.Validator, backofficeuserservice.Service, authorizationservice.Service, matchingservice.Service, matchingvalidator.Validator) {
	authSvc := authservice.New(cfg.Auth)

	MysqlRepo := mysql.New(cfg.Mysql)

	userMysql := mysqluser.New(MysqlRepo)
	userSvc := userservice.New(authSvc, userMysql)

	backofficeUserSvc := backofficeuserservice.New()

	aclMysql := mysqlaccesscontrol.New(MysqlRepo)
	authorizationSvc := authorizationservice.New(aclMysql)

	uV := uservalidator.New(userMysql)

	matchingV := matchingvalidator.New()
	redisAdapter := redis.New(cfg.Redis)
	matchingRepo := redismatching.New(redisAdapter)
	matchingSvc := matchingservice.New(cfg.MatchingService, matchingRepo)

	return authSvc, userSvc, uV, backofficeUserSvc, authorizationSvc, matchingSvc, matchingV
}
