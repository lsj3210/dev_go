package handlers

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc_server/source/rpc"
	"grpc_server/source/utils"
	"net"
	"os"
)

func _serviceRegister() {
	conf := utils.GetConf()
	serviceAddr := fmt.Sprintf("0.0.0.0:%d", conf.Server.GPort)
	fmt.Println("start grpc server on: ", serviceAddr)
	listen, err := net.Listen("tcp", serviceAddr)
	if err != nil {
		fmt.Println("start grpc server, failed to listen: %v", err)
		os.Exit(-1)
	}
	s := grpc.NewServer()
	rpc.RegisterTestServiceServer(s, &ExampleHandlers{})
	reflection.Register(s)
	if err := s.Serve(listen); err != nil {
		fmt.Println("start grpc server, failed to serve: %v", err)
		os.Exit(-1)
	}
}

func Init() {
	_serviceRegister()
}
