package config

import (
	"E-01/repository/mysql"
	"E-01/service/auth"
)

type HTTPServer struct {
	Port int
}

type Config struct {
	HTTPServer HTTPServer
	Auth       auth.Config
	Mysql      mysql.Config
}
