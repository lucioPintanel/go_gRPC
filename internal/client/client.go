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

	/*
		// Criar um novo produto
		req := &pb.CreateProdutoRequest{
			Descricao: "Novo Produto",
			Categoria: "Categoria Exemplo 3",
		}

		// Enviar a solicitação ao servidor
		r, err := client.CreateProduto(ctx, req)
		if err != nil {
			logs.Logger.Sugar().DPanicf("could not create produto: %v", err)
		}
		logs.Logger.Sugar().Debugf("Produto criado: ID=%d, Descricao=%s, DataCriacao=%s, Categoria=%s",
			r.GetId(), r.GetDescricao(), r.GetDataCriacao(), r.GetCategoria())
	*/

	req := &pb.SelectByIdRequest{
		Id: 2,
	}
	// Enviar a solicitação ao servidor
	r, err := client.SelectById(ctx, req)
	if err != nil {
		logs.Logger.Sugar().DPanicf("could not create produto: %v", err)
	}
	logs.Logger.Sugar().Debugf("Produto get: ID=%d, Descricao=%s, DataCriacao=%s, Categoria=%s",
		r.GetId(), r.GetDescricao(), r.GetDataCriacao(), r.GetCategoria())

	reqAll := &pb.SelectAllRequest{}
	// Enviar a solicitação ao servidor
	rAll, err := client.SelectAll(ctx, reqAll)
	if err != nil {
		logs.Logger.Sugar().DPanicf("could not create produto: %v", err)
	}
	for _, r := range rAll.GetProdutos() {
		logs.Logger.Sugar().Debugf("Produto r: ID=%d, Descricao=%s, DataCriacao=%s, Categoria=%s",
			r.GetId(), r.GetDescricao(), r.GetDataCriacao(), r.GetCategoria())
	}
	/*
		// Atualizar um produto existente
		reqUp := &pb.UpdateProdutoRequest{
			Id:        3, // ID do produto a ser atualizado
			Descricao: "Produto Atualizado",
			Categoria: "Nova Categoria",
		}

		// Enviar a solicitação ao servidor
		r, err = client.UpdateProduto(ctx, reqUp)
		if err != nil {
			logs.Logger.Sugar().DPanicf("could not update produto: %v", err)
		}
		logs.Logger.Sugar().Debugf("Produto atualizado: ID=%d, Descricao=%s, DataCriacao=%s, Categoria=%s",
			r.GetId(), r.GetDescricao(), r.GetDataCriacao(), r.GetCategoria())

		// Deletar um produto existente
		reqDel := &pb.DeleteProdutoRequest{
			Id: 1, // ID do produto a ser deletado
		}

		// Enviar a solicitação ao servidor
		rDel, err := client.DeleteProduto(ctx, reqDel)
		if err != nil {
			logs.Logger.Sugar().DPanicf("could not delete produto: %v", err)
		}
		if rDel.GetSucesso() {
			logs.Logger.Sugar().Debugf("Produto deletado com sucesso: ID=%d", reqDel.GetId())
		} else {
			logs.Logger.Sugar().Debugf("Falha ao deletar o produto: ID=%d", reqDel.GetId())
		} */

}
