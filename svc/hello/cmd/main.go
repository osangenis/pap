package main

import (
	"log"
	"net"
	"os"

	"github.com/osangenis/hello/src/hello"
	"google.golang.org/grpc"
)

func main() {
	// Get the port from the environment variable or use default 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create a TCP listener on the specified port
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	server := grpc.NewServer()

	// Register the HelloService implementation with the server
	hello.RegisterHelloServiceServer(server, &hello.Server{})

	// Start serving incoming requests
	log.Printf("Server started listening on port %s", port)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
