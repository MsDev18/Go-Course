package main

import (
	"E-01/repository/mysql"
	"E-01/service/userservice"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/health-check", HeathCheckHandler)
	mux.HandleFunc("/users/register", UserRegisterHandler)
	mux.HandleFunc("/users/login", UserLoginHandler)

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
	userSvc := userservice.New(mysqlRepo)

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
	userSvc := userservice.New(mysqlRepo)

	_, err = userSvc.Login(lReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error" : "%s"}`, err.Error())))
		return
	}
	writer.Write([]byte(`{"message" : "logedin credentials ok"}`))
}
