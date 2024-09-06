package main

import (
	"github.com/SidharthSasikumar/train-ticket-grpc/service"
	"github.com/SidharthSasikumar/train-ticket-grpc/ticket"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	ticket.RegisterTicketingServiceServer(grpcServer, &service.Server{
		Users: make(map[string][]*ticket.Receipt),
		Seats: make(map[string]string),
	})

	log.Printf("Server is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
