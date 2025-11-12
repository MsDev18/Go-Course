package config

import (
	"E-04/repository/mysql"
	"E-04/service/auth"
)

type HTTPServer struct {
	Port int
}

type Config struct {
	HTTPServer HTTPServer
	Auth       auth.Config
	Mysql      mysql.Config
}
