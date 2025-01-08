package server

import (
	"grpccli_srv/controllers"
	"net"
	"os"

	logs "grpccli_srv/internal/logs"
	"grpccli_srv/pb"

	"google.golang.org/grpc"
)

// Server representa o servidor gRPC.
type Server struct {
	pb.UnimplementedProdutoServiceServer // Implementa o serviço Produto.
}

// NewGRPCServer cria uma nova instância do servidor gRPC.
func NewGRPCServer() *grpc.Server {
	s := grpc.NewServer()
	return s
}

// Start inicia o servidor gRPC.
func (s *Server) Start() {
	logs.Logger.Info("Iniciando o servidor na porta de acesso 50051")

	// Configura o listener na porta 50051.
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		logs.Logger.Sugar().Errorf("Falha em inicializar listen")
		os.Exit(1)
	}
	grpcServer := NewGRPCServer()

	// Registra o controlador do produto no servidor gRPC.
	controllers.RegisterProdutoController(grpcServer)

	// Inicia o servidor gRPC.
	err = grpcServer.Serve(lis)
	if err != nil {
		logs.Logger.Sugar().Errorf("Falha em inicializar o servidor")
		os.Exit(1)
	}
	logs.Logger.Info("Finalizando o servidor")
}
