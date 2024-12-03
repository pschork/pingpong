package main

import (
	"log"
	"net"
	"net/http"

	pb "pingpong/pingpong/pkg/pingpong"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedPingServiceServer
	pb.UnimplementedPongServiceServer
}

func main() {
	port := ":50057"
	health_port := ":50067"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPingServiceServer(grpcServer, &server{})
	pb.RegisterPongServiceServer(grpcServer, &server{})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("OK")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	go http.ListenAndServe(health_port, nil) // HTTP health check server
	reflection.Register(grpcServer)

	log.Printf("Reflector is running on port %s health %s", port, health_port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
