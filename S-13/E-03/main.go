package main

import (
	"E-03/config"
	"E-03/delivery/httpserver"
	"E-03/repository/mysql"
	"E-03/service/auth"
	"E-03/service/user"
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

	authSvc, userSvc := setupServices(cfg)

	server := httpserver.New(cfg, authSvc, userSvc)

	server.Serve()
}

func setupServices(cfg config.Config) (auth.Service, user.Service) {
	MysqlRepo := mysql.New(cfg.Mysql)

	authSvc := auth.New(cfg.Auth)
	userSvc := user.New(authSvc, MysqlRepo)

	return authSvc, userSvc
}

//func UserLoginHandler(writer http.ResponseWriter, req *http.Request) {
//	writer.Header().Set("Content-Type", "application/json")
//	if req.Method != http.MethodPost {
//		fmt.Fprint(writer, `{"error" : "invalid method"}`)
//		return
//	}
//
//	data, err := io.ReadAll(req.Body)
//	if err != nil {
//		writer.Write([]byte(fmt.Sprintf(`{"error" : "%s"}`, err.Error())))
//		return
//	}
//
//	var lReq user.LoginRequest
//	err = json.Unmarshal(data, &lReq)
//	if err != nil {
//		writer.Write([]byte(fmt.Sprintf(`{"error" : "%s"}`, err.Error())))
//		return
//	}
//
//	authSvc := auth.New(JwtSignKey, AccessTokenSubject, RefreshTokenSubject, AccessTokenExpireDuration, RefreshTokenExpireDuration)
//
//	mysqlRepo := mysql.New()
//	userSvc := user.New(authSvc, mysqlRepo)
//
//	resp, err := userSvc.Login(lReq)
//	if err != nil {
//		writer.WriteHeader(http.StatusBadRequest)
//		writer.Write([]byte(fmt.Sprintf(`{"error" : "%s"}`, err.Error())))
//		return
//	}
//	data, err = json.Marshal(resp)
//	if err != nil {
//		writer.Write([]byte(fmt.Sprintf(`{"error" : "%s"}`, err.Error())))
//		return
//	}
//	writer.Write(data)
//}
//
//func UserProfileHandler(writer http.ResponseWriter, req *http.Request) {
//	if req.Method != http.MethodGet {
//		fmt.Fprint(writer, `{"error" : "invalid method"}`)
//	}
//
//	authSvc := auth.New(JwtSignKey, AccessTokenSubject, RefreshTokenSubject, AccessTokenExpireDuration, RefreshTokenExpireDuration)
//
//	authToken := req.Header.Get("Authorization")
//	claims, err := authSvc.ParseToken(authToken)
//	if err != nil {
//		fmt.Fprintf(writer, `{"error" : "token is not valid"}`)
//	}
//
//	mysqlRepo := mysql.New()
//	userSvc := user.New(authSvc, mysqlRepo)
//
//	resp, err := userSvc.Profile(user.ProfileRequest{
//		UserID: claims.UserID,
//	})
//	if err != nil {
//		writer.Write([]byte(fmt.Sprintf(`{"error" : "%s"}`, err.Error())))
//		return
//	}
//
//	data, err := json.Marshal(resp)
//	if err != nil {
//		writer.Write([]byte(fmt.Sprintf(`{"error" : "%s"}`, err.Error())))
//		return
//	}
//
//	writer.Write(data)
//}
