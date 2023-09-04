
package hello

import (
	"context"
)

type Server struct{}

func (s *Server) Ping(ctx context.Context, req *PingRequest) (*PingResponse, error) {
	return &PingResponse{
		Out: req.In,
	}, nil
}
