package attachmentservice

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"path/filepath"
	att "services/attachment"
	"services/utils"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
)

// Server server
type Server struct{}

// GetService get service
func GetService() *Server {
	return &Server{}
}

// AttachmentCreateFile attachment create file
func (s *Server) AttachmentCreateFile(ctx context.Context, in *att.AttachmentCreateFileRequest) (*att.AttachmentCreateFileResponse, error) {
	if len(in.FileName) > 0 && len(in.Muid) > 0 {
		path := os.Getenv("UPLOAD_PATH") + in.Muid
		err := os.MkdirAll(path, 0777)
		if err == nil {
			file := path + "/" + in.FileName
			err = os.Setenv("UPLOAD_FILENAME", file)
			if err == nil {
				filePath, _ := filepath.Abs(file)
				fd, err := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
				defer fd.Close()
				if err == nil {
					return &att.AttachmentCreateFileResponse{Status: codes.OK.String(), FilePath: filePath}, nil
				}
			}
		}
	}
	return &att.AttachmentCreateFileResponse{Status: codes.DataLoss.String(), FilePath: ""}, nil
}

// WriteAttachmentFile make transaction
func (s *Server) WriteAttachmentFile(stream att.AttachmentPostTxn_WriteAttachmentFileServer) error {
	str := ""
	path, _ := filepath.Abs(os.Getenv("UPLOAD_FILENAME"))
	for {
		in, err := stream.Recv()
		if err != nil {
			fileData, err := base64.URLEncoding.DecodeString(str)
			utils.Panic(err)
			ioutil.WriteFile(path, fileData, 0777)
			break
		}
		str = str + in.Data
	}
	return nil
}
