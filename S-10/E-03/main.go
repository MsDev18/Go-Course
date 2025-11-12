package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	server := HttpServer{}
	http.ListenAndServe(":8080", server)

}

type HttpServer struct {
}

func (h HttpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" && r.URL.Path == "/" {
		u := User{
			Name:  "MsDEv18",
			Age:   18,
			Phone: "0935-172-1415",
		}
		bufferdJson, err := json.Marshal(u)
		if err !=nil {
			fmt.Println("Error in Write Response ...")
			return
		}
		w.Write(bufferdJson)
	}
}

type User struct {
	Name  string
	Age   int
	Phone string
}
