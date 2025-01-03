package server

import (
	"grpccli_srv/pb"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
    pb.UnimplementedMeuServicoServer
}

func NewGRPCServer() *grpc.Server {
    certFile := "certificados/server.crt"
    keyFile := "certificados/server.key"
    creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
    if err != nil {
        os.Exit(1)
    }
    s := grpc.NewServer(grpc.Creds(creds))
    pb.RegisterMeuServicoServer(s, &server{})
    return s
}

func (s *server) Start() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        os.Exit(1)
    }
    grpcServer := NewGRPCServer()
    if err := grpcServer.Serve(lis); err != nil {
        os.Exit(1)
    }
}
