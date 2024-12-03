package main

import (
	"context"
	"log"
	"net"

	pb "pingpong/pingpong/pkg/pingpong"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type PingServer struct {
	pb.UnimplementedPingServiceServer
}

func (s *PingServer) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	log.Printf("Ping received: %s", req.Message)
	return &pb.PingResponse{Reply: "Ping from PingService"}, nil
}

func main() {
	port := "50051"
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPingServiceServer(grpcServer, &PingServer{})

	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)

	// Set the health status to "SERVING"
	healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)

	reflection.Register(grpcServer)

	log.Printf("PingService is running with health on port %s...", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
