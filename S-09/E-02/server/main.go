package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	//fmt.Println("-----<<< Say Hi To Go >>>-----")
	const (
		network = "tcp"
		address = "127.0.0.1:80"
	)
	// Create New Listener
	listener, err := net.Listen(network, address)
	if err != nil {
		log.Fatalln("can't listen on given address : ", address, err)
	}
	fmt.Println("Listening on address : ", listener.Addr())
	// Listen For New Connections
	connection, aErr := listener.Accept()
	if aErr != nil {
		log.Fatalln("can't listen to new connection : ", aErr)
	}
	fmt.Println("\nServer : ")
	fmt.Println("Local Address : ", connection.LocalAddr())
	fmt.Println("Remote Address (Client) : " , connection.RemoteAddr())
	// Process Request
}
