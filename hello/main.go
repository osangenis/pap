
package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/osangenis/hello/hello"
	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	hello.RegisterHelloServiceServer(s, &hello.Server{})

	log.Printf("Server listening on port %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
