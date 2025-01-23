package server

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/xIndustries/BandRoom/backend-auth/proto/Generated"
)

// RunGRPCServer starts the gRPC server.
func RunGRPCServer(port string, handler pb.UserServiceServer) error {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, handler)

	// Enable gRPC Reflection
	reflection.Register(server)

	log.Printf("gRPC server is listening on port %s", port)
	return server.Serve(listener)
}
