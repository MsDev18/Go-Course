package main

import (
	"fmt"
	"net"
)

func main() {
	var response = `HTTP/1.1 500 OK
Date: Sat, 09 Oct 2010 14:28:02 GMT
Server: GO Server
Last-Modified: Tue, 01 Dec 2009 20:18:22 GMT
ETag: "51142bc1-7449-479b075b2891b"
Accept-Ranges: bytes
Content-Length: 4
Content-Type: text/html

TEST`


	listener, _ := net.Listen("tcp", ":8080")
	conn, err := listener.Accept()

	if err != nil {
		fmt.Println("Error in Accept connection", err)

		return
	}

	var data = make([]byte, 1024)
	_, rErr := conn.Read(data)

	if rErr != nil {
		fmt.Println("Error in read : ", rErr)

		return
	}

	fmt.Println(string(data))

	_, wErr := conn.Write([]byte(response))

	if wErr != nil {
		fmt.Println("Error in Write : ", wErr)

		return
	}
	conn.Close()
}
