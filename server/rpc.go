package server

import (
	"fmt"
	"net"

	pb "grpc-gw/example"
	"grpc-gw/handlers"

	"google.golang.org/grpc"
)

func RunRPC() error {
	fmt.Println("rpc called")

	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		return err
	}

	defer func() {
		err = lis.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	rpcServer := grpc.NewServer()
	mysrv := &handlers.RpcServer{}
	pb.RegisterYourServiceServer(rpcServer, mysrv)
	return rpcServer.Serve(lis)
}
