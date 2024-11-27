package main

import (
	"log"
	"net"

	pb "pingpong/pingpong/pkg/pingpong"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedPingServiceServer
	pb.UnimplementedPongServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":50057")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPingServiceServer(s, &server{})
	pb.RegisterPongServiceServer(s, &server{})
	reflection.Register(s)

	log.Println("Reflector is running on port :50057")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
