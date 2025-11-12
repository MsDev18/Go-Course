package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	message := "Default Message"
	if len(os.Args) > 1 {
		message = os.Args[1]
	}
	connection, err := net.Dial("tcp", "127.0.0.1:80")
	if err != nil {
		log.Fatalln("Can't Dial The Given Address : ", err)
	}
	
	fmt.Println("Local Address : " ,connection.LocalAddr())

	numberOfWrittenBytes,wErr := connection.Write([]byte(message))
	if wErr != nil {
		log.Fatalln("Can't Write Data To Connection" , wErr)
	}
	fmt.Println("Number of Written Bytes : " , numberOfWrittenBytes)

	data := make([]byte , 1024)
	_, rErr := connection.Read(data)
	if rErr != nil {
		log.Fatalln(`Can't Read Data From Connection : `, rErr)
	}

	fmt.Println("Server Response : ", string(data))
}