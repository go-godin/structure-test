package main

import (
	"github.com/go-godin/example-structure/internal/endpoint"
	"github.com/go-godin/example-structure/internal/ticket"
	"github.com/go-godin/log"
	pb "bitbucket.org/jdbergmann/protobuf-go/ticket/ticket"
	googleGrpc "google.golang.org/grpc"
	"net"
)

func main() {

	logger := log.NewLoggerFromEnv()

	var service ticket.Service
	service = ticket.NewService(logger)

	endpoints := endpoint.NewEndpointSet(service)

	grpcServer := googleGrpc.NewServer()
	grpcListener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
	    panic(err)
	}
	pb.RegisterTicketServiceServer(grpcServer, endpoints)

	if err := grpcServer.Serve(grpcListener); err != nil {
	   panic(err)
	}

}
