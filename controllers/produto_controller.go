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

func (c *ProdutoController) UpdateProduto(ctx context.Context, req *pb.UpdateProdutoRequest) (*pb.ProdutoResponse, error) {
	// Buscar o produto existente no banco de dados para obter a DataCriacao original
	produtoFromDB, err := repositories.GetProdutoByID(int(req.GetId()))
	if err != nil {
		return nil, err
	}

	// Criar o objeto Produto com a DataCriacao preservada
	produto := models.Produto{
		ID:          produtoFromDB.ID,
		Descricao:   req.GetDescricao(),
		DataCriacao: produtoFromDB.DataCriacao,
		Categoria:   req.GetCategoria(),
	}

	// Atualizar o produto no banco de dados
	err = repositories.UpdateProduto(produto)
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
