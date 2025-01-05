package client

import (
	"context"
	"grpccli_srv/pb"
	"time"

	logs "grpccli_srv/internal/logs"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

func RunClient() {
	// Estabelecer a conexão com o servidor gRPC sem TLS
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logs.Logger.Sugar().DPanicf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewProdutoServiceClient(conn)

	// Preparar o contexto com timeout para a solicitação
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Criar um novo produto
	req := &pb.CreateProdutoRequest{
		Descricao: "Novo Produto",
		Categoria: "Categoria Exemplo",
	}

	// Enviar a solicitação ao servidor
	r, err := client.CreateProduto(ctx, req)
	if err != nil {
		logs.Logger.Sugar().DPanicf("could not create produto: %v", err)
	}
	logs.Logger.Sugar().Debugf("Produto criado: ID=%d, Descricao=%s, DataCriacao=%s, Categoria=%s",
		r.GetId(), r.GetDescricao(), r.GetDataCriacao(), r.GetCategoria())
}
