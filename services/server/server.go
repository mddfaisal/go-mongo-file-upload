package server

import (
	"net"
	"os"
	att "services/attachment"
	"services/attachmentservice"
	eb "services/emailbody"
	"services/emailbodyservice"
	"services/utils"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Run run server
func Run() {
	listen, err := net.Listen("tcp", os.Getenv("RPC_PORT"))
	utils.Panic(err)
	ser := grpc.NewServer()
	eb.RegisterEmailBodyPostTxnServer(ser, emailbodyservice.GetService())
	att.RegisterAttachmentCreateFileServer(ser, attachmentservice.GetService())
	att.RegisterAttachmentPostTxnServer(ser, attachmentservice.GetService())
	reflection.Register(ser)
	if err := ser.Serve(listen); err != nil {
		utils.Panic(err)
	}
}
