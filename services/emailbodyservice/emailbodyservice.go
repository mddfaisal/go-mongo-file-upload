package emailbodyservice

import (
	"services/email"
	eb "services/emailbody"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
)

// Server server
type Server struct{}

// GetService get email body service
func GetService() *Server {
	return &Server{}
}

// GetEmailBody get email
func (s *Server) GetEmailBody(ctx context.Context, in *eb.EmailBodyGetRequest) (*eb.EmailBodyGetResponse, error) {
	if len(in.Id) == 0 {
		return &eb.EmailBodyGetResponse{Id: "", Muid: "", Subject: "", Html: "", Status: codes.DataLoss.String()}, nil
	}
	e := email.Email{ID: in.Id}
	err := e.FindOne()
	if err != nil {
		return &eb.EmailBodyGetResponse{Id: "", Muid: "", Subject: "", Html: "", Status: codes.NotFound.String()}, nil
	}
	return &eb.EmailBodyGetResponse{
		Id:      e.ID,
		Muid:    e.Muid,
		Subject: e.Subject,
		Html:    e.HTML,
	}, nil
}

// PostEmailBody make transaction
func (s *Server) PostEmailBody(ctx context.Context, in *eb.EmailBodyPostRequest) (*eb.EmailBodyPostResponse, error) {
	if len(in.Html) == 0 || len(in.Muid) == 0 || len(in.Subject) == 0 {
		return &eb.EmailBodyPostResponse{Id: "", Status: codes.DataLoss.String()}, nil
	}
	e := email.Email{
		Muid:    in.Muid,
		HTML:    in.Html,
		Subject: in.Subject,
	}
	e.DecodeHTML()
	if len(e.HTML) == 0 {
		return &eb.EmailBodyPostResponse{Id: "", Status: codes.DataLoss.String()}, nil
	}
	id := e.Create()
	return &eb.EmailBodyPostResponse{Id: id, Status: codes.OK.String()}, nil
}
