package main

import "E-01/delivery/grpcserver/presenceserver"

func main() {
	server := presenceserver.Server{}
	server.Start()

}