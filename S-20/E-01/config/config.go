package config

import (
	"E-01/adapter/redis"
	"E-01/repository/mysql"
	"E-01/service/authservice"
	"E-01/service/matchingservice"
	"time"
)

type Application struct {
	GracefulShotDownTimeout time.Duration `json:"graceful_shotdown_timeout"`
}

type HTTPServer struct {
	Port int `koanf:"port"`
}

type Config struct {
	Application     Application            `koanf:"application"`
	HTTPServer      HTTPServer             `koanf:"http_server"`
	Auth            authservice.Config     `koanf:"auth"`
	Mysql           mysql.Config           `koanf:"mysql"`
	MatchingService matchingservice.Config `koanf:"matching_service"`
	Redis           redis.Config           `koanf:"redis"`
}
