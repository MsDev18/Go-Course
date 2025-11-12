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

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users/register", UserRegisterHandler)
	mux.HandleFunc("/health-check", HeathCheckHandler)

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
