package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	connection, err := net.Dial("tcp", "127.0.0.1:80")
	if err != nil {
		log.Fatalln("can't dial the given address : ", err)
	}
	fmt.Println("\nClient : ")
	fmt.Println("Local Address : ", connection.LocalAddr())
	fmt.Println("Remote Address (Server) : ", connection.RemoteAddr())
	connection2, err := net.Dial("tcp", "127.0.0.1:80")
	if err != nil {
		log.Fatalln("can't dial the given address : ", err)
	}
	fmt.Println("\nClient : ")
	fmt.Println("Local Address : ", connection2.LocalAddr())
	fmt.Println("Remote Address (Server) : ", connection2.RemoteAddr())
}

// Watched To S-09 ||| E-02 ||| 35:00 Minutes
