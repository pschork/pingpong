package main

import (
	"log"
	"net"

	pb "pingpong/pingpong/pkg/pingpong"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
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
	grpcServer := grpc.NewServer()
	pb.RegisterPingServiceServer(grpcServer, &server{})
	pb.RegisterPongServiceServer(grpcServer, &server{})

	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)
	healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)

	reflection.Register(grpcServer)

	log.Println("Reflector is running with health on port :50057")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
