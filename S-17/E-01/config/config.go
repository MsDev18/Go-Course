package config

import (
	"E-01/repository/mysql"
	"E-01/service/authservice"
)

type HTTPServer struct {
	Port int
}

type Config struct {
	HTTPServer HTTPServer
	Auth       authservice.Config
	Mysql      mysql.Config
}
