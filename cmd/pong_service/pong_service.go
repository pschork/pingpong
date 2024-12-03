package main

import (
	"context"
	"log"
	"net"
	"net/http"

	pb "pingpong/pingpong/pkg/pingpong"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
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
	port := ":50052"
	health_port := ":50062"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPongServiceServer(grpcServer, &PongServer{})
	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("OK")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	go http.ListenAndServe(health_port, nil) // HTTP health check server

	reflection.Register(grpcServer)

	log.Printf("PongService is running on port %s health %s", port, health_port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
