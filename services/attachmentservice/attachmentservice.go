package attachmentservice

import (
	"fmt"
	att "services/attachment"

	"golang.org/x/net/context"
)

// Server server
type Server struct{}

// GetService get service
func GetService() *Server {
	return &Server{}
}

// PostAttachment make transaction
func (s *Server) PostAttachment(ctx context.Context, in *att.AttachmentPostRequest) (*att.AttachmentPostResponse, error) {
	fmt.Println(in.Muid)
	fmt.Println(in.Filename)
	fmt.Println(in.Data)
	return &att.AttachmentPostResponse{}, nil
}
