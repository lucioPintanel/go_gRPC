package controllers

import (
	"context"
	"grpccli_srv/models"
	"grpccli_srv/pb"
	"grpccli_srv/repositories"
	"time"

	"google.golang.org/grpc"
)

// ProdutoController implementa o servidor gRPC para o serviço Produto.
type ProdutoController struct {
	pb.UnimplementedProdutoServiceServer
}

// CreateProduto cria um novo produto com base nas informações fornecidas no request.
func (c *ProdutoController) CreateProduto(ctx context.Context, req *pb.CreateProdutoRequest) (*pb.ProdutoResponse, error) {
	// Cria uma nova instância de Produto a partir do request.
	produto := models.Produto{
		Descricao:   req.GetDescricao(),
		DataCriacao: time.Now(),
		Categoria:   req.GetCategoria(),
	}

	// Insere o novo produto no banco de dados através do repositório.
	err := repositories.CreateProduto(produto)
	if err != nil {
		return nil, err
	}

	// Retorna a resposta contendo os dados do produto criado.
	return &pb.ProdutoResponse{
		Id:          int32(produto.ID),
		Descricao:   produto.Descricao,
		DataCriacao: produto.DataCriacao.Format(time.RFC3339),
		Categoria:   produto.Categoria,
	}, nil
}

// UpdateProduto atualiza as informações de um produto existente com base no request.
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

	// Retorna a resposta contendo os dados do produto atualizado.
	return &pb.ProdutoResponse{
		Id:          int32(produto.ID),
		Descricao:   produto.Descricao,
		DataCriacao: produto.DataCriacao.Format(time.RFC3339),
		Categoria:   produto.Categoria,
	}, nil
}

// DeleteProduto remove um produto existente com base no ID fornecido no request.
func (c *ProdutoController) DeleteProduto(ctx context.Context, req *pb.DeleteProdutoRequest) (*pb.DeleteProdutoResponse, error) {
	err := repositories.DeleteProduto(int(req.GetId()))
	if err != nil {
		return nil, err
	}

	return &pb.DeleteProdutoResponse{
		Sucesso: true,
	}, nil
}

// SelectById seleciona um produto pelo ID.
func (c *ProdutoController) SelectById(ctx context.Context, req *pb.SelectByIdRequest) (*pb.ProdutoResponse, error) {
	// Recupera o produto pelo ID a partir do repositório.
	produto, err := repositories.GetProdutoByID(int(req.GetId()))
	if err != nil {
		return nil, err
	}

	// Retorna a resposta do produto no formato protobuf.
	return &pb.ProdutoResponse{
		Id:          int32(produto.ID),
		Descricao:   produto.Descricao,
		DataCriacao: produto.DataCriacao.Format(time.RFC3339),
		Categoria:   produto.Categoria,
	}, nil
}

// SelectAll seleciona todos os produtos.
func (c *ProdutoController) SelectAll(ctx context.Context, req *pb.SelectAllRequest) (*pb.SelectAllResponse, error) {
	// Recupera todos os produtos do repositório.
	produtos, err := repositories.GetAllProdutos()
	if err != nil {
		return nil, err
	}

	// Converte a lista de produtos para a resposta protobuf.
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

	// Retorna a lista de produtos no formato protobuf.
	return &pb.SelectAllResponse{
		Produtos: produtoResponses,
	}, nil
}

// RegisterProdutoController registra o ProdutoController no servidor gRPC.
func RegisterProdutoController(server *grpc.Server) {
	pb.RegisterProdutoServiceServer(server, &ProdutoController{})
}
