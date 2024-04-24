package main

import (
	"log"
	"net"

	pb "repository_lock/protobufs"
	"repository_lock/services"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRepositoryLockServer(s, &services.RepositoryLockService{})
	log.Println("Server started on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
