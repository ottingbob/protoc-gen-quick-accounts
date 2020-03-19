package server

import (
	bb "grpc-gw/buffb"
	"grpc-gw/handlers"

	"fmt"
	"net"

	"google.golang.org/grpc"
)

func RunQuickAccountRPC() error {
	fmt.Println("qa rpc called")

	lis, err := net.Listen("tcp", ":5002")
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
	qaController := handlers.CreateQAController()
	qsServer := &handlers.QuickRPCServer{
		Name:         "Rock-Steady",
		QAController: qaController,
	}
	bb.RegisterQuickServiceServer(rpcServer, qsServer)
	return rpcServer.Serve(lis)
}
