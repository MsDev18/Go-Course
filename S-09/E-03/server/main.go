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
		log.Fatalln("Can't Listen On Given Address : ", address, err)
	}
	fmt.Println("Listening On : ", listener.Addr())

	// Listen For New Connections
	for {
		connection, aErr := listener.Accept()
		if aErr != nil {
			log.Println("Can't Listen To New Connection : ", aErr)

			continue
		}

		// Process Request
		buffer := make([]byte, 1024)
		numberOfReadBytes, rErr := connection.Read(buffer)
		if rErr != nil {
			log.Println("Cant't Read The Request ...")

			continue
		}
		fmt.Printf("Client Address : %s , numverOfReadBytes %d , Data : %s \n" , connection.RemoteAddr(),numberOfReadBytes , string(buffer))
	
		_, wErr := connection.Write([]byte(`Your Message Received`))
		if wErr != nil {
			log.Println("Cant't Write The Request ...")

			continue
		}
	}
}
