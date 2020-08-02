package emailbodyservice

import (
	"fmt"
	eb "services/emailbody"

	"golang.org/x/net/context"
)

// Server server
type Server struct{}

// GetService get email body service
func GetService() *Server {
	return &Server{}
}

// MakeTransaction make transaction
func (s *Server) MakeTransaction(ctx context.Context, in *eb.EmailBodyRequest) (*eb.EmailBodyResponse, error) {
	fmt.Println(in.Muid)
	fmt.Println(in.Subject)
	return &eb.EmailBodyResponse{Id: "hello world..."}, nil
}
