package server

import (
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func RunGRPCServer(addr string, registerServer func(*grpc.Server)) {
	grpcServer := grpc.NewServer()
	registerServer(grpcServer)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.Panicf("failed to listen at %s: %v", addr, err)
	}

	logrus.WithField("gRPC endpoint", addr).Info("Starting gRPC Listener")
	logrus.Fatal(grpcServer.Serve(lis))
}
