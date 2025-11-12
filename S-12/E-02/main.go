package main

import (
	"E-02/repository/mysql"
	"E-02/service/userservice"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	JwtSignKey = "jwt_secret"
)

func main() {

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

	var uReq userservice.RegisterRequest
	err = json.Unmarshal(data, &uReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error" : "%s"}`, err.Error())))
		return
	}
	mysqlRepo := mysql.New()
	userSvc := userservice.New(mysqlRepo, JwtSignKey)

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

	var lReq userservice.LoginRequest
	err = json.Unmarshal(data, &lReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error" : "%s"}`, err.Error())))
		return
	}
	mysqlRepo := mysql.New()
	userSvc := userservice.New(mysqlRepo, JwtSignKey)

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
	
	
	pReq := userservice.ProfileRequest{UserID: 0}
	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error" : "%s"}`, err.Error())))
		return
	}

	err = json.Unmarshal(data, &pReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error" : "%s"}`, err.Error())))
		return
	}
	mysqlRepo := mysql.New()
	userSvc := userservice.New(mysqlRepo, JwtSignKey)

	resp, err := userSvc.Profile(pReq)
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
