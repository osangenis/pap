
package hello

import (
	"context"
)

// Server represents the server implementation for HelloService.
type Server struct{}

// Ping implements the Ping method of the HelloService.
func (s *Server) Ping(ctx context.Context, req *PingRequest) (*PingResponse, error) {
	// Create a PingResponse with the same string as the PingRequest
	return &PingResponse{Out: req.In}, nil
}
