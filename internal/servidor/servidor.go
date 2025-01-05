package server

import (
	"grpccli_srv/controllers"
	"net"
	"os"

	logs "grpccli_srv/internal/logs"
	"grpccli_srv/pb"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedProdutoServiceServer
}

func NewGRPCServer() *grpc.Server {
	s := grpc.NewServer()
	return s
}

func (s *Server) Start() {
	logs.Logger.Info("Iniciando o servidor na porta de acesso 50051")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		logs.Logger.Sugar().Errorf("Falha em inicializar listen")
		os.Exit(1)
	}
	grpcServer := NewGRPCServer()

	// Registrar o controlador do produto apenas aqui
	controllers.RegisterProdutoController(grpcServer)

	err = grpcServer.Serve(lis)
	if err != nil {
		logs.Logger.Sugar().Errorf("Falha em inicializar o servidor")
		os.Exit(1)
	}
	logs.Logger.Info("Finalizando o servidor")
}
