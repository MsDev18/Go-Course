package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	var response = []byte(`HTTP/1.1 500 OK
Date: Sat, 09 Oct 2010 14:28:02 GMT
Server: Apache
Last-Modified: Tue, 01 Dec 2009 20:18:22 GMT
ETag: "51142bc1-7449-479b075b2891b"
Accept-Ranges: bytes
Content-Length: 29769
Content-Type: text/html

Test Messaeg From Server`)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening : ", err)
		return
	}
	//fmt.Println(string(wData))

	conn, cErr := listener.Accept()
	if cErr != nil {
		fmt.Println("Error accepting : ", cErr)
		return
	}
	defer conn.Close()

	var request = make([]byte, 1024)
	_, rErr := conn.Read(request)
	if rErr != nil {
		fmt.Println("Error reading : ", rErr)
		return
	}
	fmt.Println(string(request))
	method, path := ParseHttpRequest(string(request))
	if method == "GET" && path != "/" {
		var fData = make([]byte, 20480)
		f, _ := os.OpenFile(strings.Replace(path, "/", "", 1)+".html", os.O_RDONLY, 0777)
		fData, _ = io.ReadAll(f)
		contentLength := len(fData)
		contentType := "text/html"
		responseStatusCode := 200
		responseStatusMessage := "OK"

		response = CreateHttpResponse(fData, contentType, responseStatusMessage, contentLength, responseStatusCode)
	}
	fmt.Println(string(response))
	_, wErr := conn.Write(response)
	if wErr != nil {
		fmt.Println("Error writing : ", wErr)
		return
	}
}

func ParseHttpRequest(request string) (string, string) {
	lines := strings.Split(request, "\n")

	req := strings.Split(lines[0], " ")
	method, path, httpVersion := req[0], req[1], req[2]
	var host string
	for _, line := range lines {
		if strings.Contains(line, "Host") {
			host = strings.Replace(line, "Host: ", "", 1)
		}
	}
	fmt.Println("-----<<< Started >>>-----")
	fmt.Println("Method : ", method)
	fmt.Println("Host : ", host)
	fmt.Println("Path : ", path)
	fmt.Println("http version : ", httpVersion)
	fmt.Println("-----<<< Finished >>------")
	return method, path
}

func CreateHttpResponse(fData []byte, contentType, responseStatusMessage string, contentLength, responseStatusCode int) []byte {
	return []byte(fmt.Sprintf(`HTTP/1.1 %d %s
Content-Length : %d
Content-Type: %s

%s`, responseStatusCode, responseStatusMessage, contentLength, contentType, fData))
}
