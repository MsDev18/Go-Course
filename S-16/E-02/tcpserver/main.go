package main

import (
	"log"
	"net"
)

func main() {
	server , err := net.Listen("tcp" , ":8099")
	if err != nil {
		panic(err)
	}

	for {
		connection, err := server.Accept()
		if err != nil {
			panic(err)
		}
		log.Println("connection.RemoteAddr() ",connection.RemoteAddr())
		
		buffer := make([]byte, 1024)
		_, err = connection.Read(buffer)
		if err !=nil {
			panic(err)
		}

		log.Println("data : ", string(buffer))
	}
	server.Close()
}