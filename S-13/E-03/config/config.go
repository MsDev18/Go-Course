package config

import (
	"E-03/repository/mysql"
	"E-03/service/auth"
)

type HTTPServer struct {
	Port int
}

type Config struct {
	HTTPServer HTTPServer
	Auth       auth.Config
	Mysql      mysql.Config
}
