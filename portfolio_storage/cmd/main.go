package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	pb "stocks/grpc_utils/proto/PortfolioStorage"
	"stocks/portfolio_storage/internal/server"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	server := server.CreateServer()
	pb.RegisterPortfolioStorageServer(grpcServer, server)
	grpcServer.Serve(lis)
}