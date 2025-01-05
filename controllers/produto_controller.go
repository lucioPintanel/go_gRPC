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

func (c *ProdutoController) DeleteProduto(ctx context.Context, req *pb.DeleteProdutoRequest) (*pb.DeleteProdutoResponse, error) {
	err := repositories.DeleteProduto(int(req.GetId()))
	if err != nil {
		return nil, err
	}

	return &pb.DeleteProdutoResponse{
		Sucesso: true,
	}, nil
}

func (c *ProdutoController) SelectById(ctx context.Context, req *pb.SelectByIdRequest) (*pb.ProdutoResponse, error) {
	produto, err := repositories.GetProdutoByID(int(req.GetId()))
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

func (c *ProdutoController) SelectAll(ctx context.Context, req *pb.SelectAllRequest) (*pb.SelectAllResponse, error) {
	produtos, err := repositories.GetAllProdutos()
	if err != nil {
		return nil, err
	}

	var produtoResponses []*pb.ProdutoResponse
	for _, produto := range produtos {
		produtoResponse := &pb.ProdutoResponse{
			Id:          int32(produto.ID),
			Descricao:   produto.Descricao,
			DataCriacao: produto.DataCriacao.Format(time.RFC3339),
			Categoria:   produto.Categoria,
		}
		produtoResponses = append(produtoResponses, produtoResponse)
	}

	return &pb.SelectAllResponse{
		Produtos: produtoResponses,
	}, nil
}

func RegisterProdutoController(server *grpc.Server) {
	pb.RegisterProdutoServiceServer(server, &ProdutoController{})
}
