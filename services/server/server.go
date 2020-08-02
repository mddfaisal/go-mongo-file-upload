package server

import (
	"fmt"
	"net"
	"os"
	att "services/attachment"
	"services/attachmentservice"
	eb "services/emailbody"
	"services/emailbodyservice"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Run run server
func Run() {
	listen, err := net.Listen("tcp", os.Getenv("RPC_PORT"))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	ser := grpc.NewServer()
	eb.RegisterEmailBodyTransactionServer(ser, emailbodyservice.GetService())
	att.RegisterAttachmentTransactionServer(ser, attachmentservice.GetService())
	reflection.Register(ser)
	if err := ser.Serve(listen); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
