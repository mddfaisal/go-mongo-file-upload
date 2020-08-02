package attachmentservice

import (
	att "services/attachment"

	"golang.org/x/net/context"
)

// Server server
type Server struct{}

// GetService get service
func GetService() *Server {
	return &Server{}
}

// MakeTransaction make transaction
func (s *Server) MakeTransaction(ctx context.Context, in *att.AttachmentRequest) (*att.AttachmentResponse, error) {
	return &att.AttachmentResponse{}, nil
}
