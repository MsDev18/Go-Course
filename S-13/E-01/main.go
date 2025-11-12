package main

import (
	"E-01/repository/mysql"
	"E-01/service/auth"
	"E-01/service/user"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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
	fmt.
	mux := http.NewServeMux()
	mux.HandleFunc("/health-check", HeathCheckHandler)
	mux.HandleFunc("/users/register", UserRegisterHandler)
	mux.HandleFunc("/users/login", UserLoginHandler)
	mux.HandleFunc("/users/profile", UserProfileHandler)

	log.Println("Server Listening On Port 8080 ...")
	server := http.Server{Addr: ":8080", Handler: mux}
	log.Fatal(server.ListenAndServe())

}

func HeathCheckHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, `{"message" : "Every Think Is Good !"}`)
}

func UserRegisterHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		fmt.Fprint(writer, `{"error" : "invalid method"}`)
		return
	}

	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error" : "%s"}`, err.Error())))
		return
	}

	var uReq user.RegisterRequest
	err = json.Unmarshal(data, &uReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error" : "%s"}`, err.Error())))
		return
	}

	authSvc := auth.New(JwtSignKey, AccessTokenSubject, RefreshTokenSubject, AccessTokenExpireDuration, RefreshTokenExpireDuration)

	mysqlRepo := mysql.New()
	userSvc := user.New(authSvc, mysqlRepo)

	_, err = userSvc.Register(uReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error" : "%s"}`, err.Error())))
		return
	}
	writer.Write([]byte(`{"message" : "User Created"}`))
}

func UserLoginHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		fmt.Fprint(writer, `{"error" : "invalid method"}`)
		return
	}

	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error" : "%s"}`, err.Error())))
		return
	}

	var lReq user.LoginRequest
	err = json.Unmarshal(data, &lReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error" : "%s"}`, err.Error())))
		return
	}

	authSvc := auth.New(JwtSignKey, AccessTokenSubject, RefreshTokenSubject, AccessTokenExpireDuration, RefreshTokenExpireDuration)

	mysqlRepo := mysql.New()
	userSvc := user.New(authSvc, mysqlRepo)

	resp, err := userSvc.Login(lReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error" : "%s"}`, err.Error())))
		return
	}
	data, err = json.Marshal(resp)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error" : "%s"}`, err.Error())))
		return
	}
	writer.Write(data)
}

func UserProfileHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		fmt.Fprint(writer, `{"error" : "invalid method"}`)
	}

	authSvc := auth.New(JwtSignKey, AccessTokenSubject, RefreshTokenSubject, AccessTokenExpireDuration, RefreshTokenExpireDuration)

	authToken := req.Header.Get("Authorization")
	claims, err := authSvc.ParseToken(authToken)
	if err != nil {
		fmt.Fprintf(writer, `{"error" : "token is not valid"}`)
	}

	mysqlRepo := mysql.New()
	userSvc := user.New(authSvc, mysqlRepo)

	resp, err := userSvc.Profile(user.ProfileRequest{
		UserID: claims.UserID,
	})
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error" : "%s"}`, err.Error())))
		return
	}

	data, err := json.Marshal(resp)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error" : "%s"}`, err.Error())))
		return
	}

	writer.Write(data)
}
