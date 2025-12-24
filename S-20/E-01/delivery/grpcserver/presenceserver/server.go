package presenceserver

import (
	"E-01/contract/golang/presence"
	"E-01/param"
	"E-01/service/presenceservice"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	presence.UnimplementedPresenceServiceServer
	svc presenceservice.Service
}

func (s Server) GetPresence(ctx context.Context,req *presence.GetPresenceRequest) (*presence.GetPresenceResponse, error) {
	s.svc.GetPresence(ctx, param.GetPresenceRequest{UserIDs: req.GetUserIds()})
}

func (s Server) Start() {
	//listenser := tcp port
	address := fmt.Sprintf(":%d", 8086)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	// pb-presenceserver
	presenceSvcServer := Server{}
	// presenceSvcServer.Start()
	//grpc-server
	grpcServer := grpc.NewServer()
	//pb-presenceserver register into grpcserver
	presence.RegisterPresenceServiceServer(grpcServer, &presenceSvcServer)
	// server grpc server by listener
	log.Printf("presence grpc server started on %s\n", address)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("couldn't srever presence grpc")
	}
}
