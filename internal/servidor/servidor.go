package server

import (
	"context"
	"net"
	"os"

	logs "grpccli_srv/internal/logs"
	"grpccli_srv/pb"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedMeuServicoServer
}

func NewGRPCServer() *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterMeuServicoServer(s, &Server{})
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
	err = grpcServer.Serve(lis)
	if err != nil {
		logs.Logger.Sugar().Errorf("Falha em inicializar o servidor")
		os.Exit(1)
	}
	logs.Logger.Info("Finilizando o servidor")
}

func (s *Server) MeuMetodo(ctx context.Context, req *pb.MeuRequest) (*pb.MeuResponse, error) {
	logs.Logger.Sugar().Infof("Recebido pedido: %v", req.Mensagem)
	return &pb.MeuResponse{Resposta: "Mensagem recebida: " + req.Mensagem}, nil
}
