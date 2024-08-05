package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func RegisterGrpc(port int64, server func(s *grpc.Server)) error {
	//声明了tcp链接
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	s := grpc.NewServer()

	server(s)

	//反射打开后可以查询API
	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		return err
	}
	return err
}
