package controllers

import (
	"context"
	"grpccli_srv/models"
	"grpccli_srv/pb"
	"grpccli_srv/repositories"
	"time"

	"google.golang.org/grpc"
)

type ProdutoController struct {
	pb.UnimplementedProdutoServiceServer
}

func (c *ProdutoController) CreateProduto(ctx context.Context, req *pb.CreateProdutoRequest) (*pb.ProdutoResponse, error) {
	produto := models.Produto{
		Descricao:   req.GetDescricao(),
		DataCriacao: time.Now(),
		Categoria:   req.GetCategoria(),
	}

	err := repositories.CreateProduto(produto)
	if err != nil {
		return nil, err
	}

	return &pb.ProdutoResponse{
		Id:          int32(produto.ID),
		Descricao:   produto.Descricao,
		DataCriacao: produto.DataCriacao.Format(time.RFC3339),
		Categoria:   produto.Categoria,
	}, nil
}

func RegisterProdutoController(server *grpc.Server) {
	pb.RegisterProdutoServiceServer(server, &ProdutoController{})
}
