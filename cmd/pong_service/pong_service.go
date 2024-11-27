package main

import (
	"context"
	"log"
	"net"

	pb "pingpong/pingpong/pkg/pingpong"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type PongServer struct {
	pb.UnimplementedPongServiceServer
}

func (s *PongServer) Pong(ctx context.Context, req *pb.PongRequest) (*pb.PongResponse, error) {
	log.Printf("Pong received: %s", req.Message)
	return &pb.PongResponse{Reply: "Pong from PongService"}, nil
}

func main() {
	port := "50052"

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPongServiceServer(grpcServer, &PongServer{})

	reflection.Register(grpcServer)

	log.Printf("PongService is running on port %s...", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}