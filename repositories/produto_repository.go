package repositories

import (
	"grpccli_srv/internal/configs"
	logs "grpccli_srv/internal/logs"
	"grpccli_srv/models"
)

// CreateProduto insere um novo produto no banco de dados.
func CreateProduto(produto models.Produto) error {
	query := `
    INSERT INTO produto (Descricao, DataCriacao, Categoria) 
    VALUES (?, ?, ?)
    `
	// Executa a query para inserir o produto.
	_, err := configs.DB.Exec(query, produto.Descricao, produto.DataCriacao, produto.Categoria)
	if err != nil {
		// Log de erro caso a execução da query falhe.
		logs.Logger.Error("falha em executar a query")
		return err
	}
	return nil
}

// GetProdutoByID recupera um produto pelo ID.
func GetProdutoByID(id int) (models.Produto, error) {
	query := "SELECT ID, Descricao, DataCriacao, Categoria FROM produto WHERE ID = ?"
	row := configs.DB.QueryRow(query, id)

	var produto models.Produto
	// Faz o scan do resultado da query para a estrutura Produto.
	err := row.Scan(&produto.ID, &produto.Descricao, &produto.DataCriacao, &produto.Categoria)
	if err != nil {
		// Log de erro caso a recuperação do produto falhe.
		logs.Logger.Sugar().Errorf("falha em obter o produto - ID: %d", id)
		return produto, err
	}

	return produto, nil
}

// UpdateProduto atualiza um produto existente no banco de dados.
func UpdateProduto(produto models.Produto) error {
	query := `
    UPDATE produto 
    SET Descricao = ?, DataCriacao = ?, Categoria = ? 
    WHERE ID = ?
    `
	// Executa a query para atualizar o produto.
	_, err := configs.DB.Exec(query, produto.Descricao, produto.DataCriacao, produto.Categoria, produto.ID)
	if err != nil {
		// Log de erro caso a execução da query falhe.
		logs.Logger.Sugar().Errorf("falha em fazer update - ID: %d", produto.ID)
		return err
	}
	return nil
}

// DeleteProduto deleta um produto pelo ID.
func DeleteProduto(id int) error {
	query := "DELETE FROM produto WHERE ID = ?"
	// Executa a query para deletar o produto.
	_, err := configs.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

// GetAllProdutos recupera todos os produtos do banco de dados.
func GetAllProdutos() ([]models.Produto, error) {
	query := "SELECT ID, Descricao, DataCriacao, Categoria FROM produto"
	rows, err := configs.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var produtos []models.Produto
	// Faz o scan de cada linha do resultado da query para a estrutura Produto.
	for rows.Next() {
		var produto models.Produto
		if err := rows.Scan(&produto.ID, &produto.Descricao, &produto.DataCriacao, &produto.Categoria); err != nil {
			return nil, err
		}
		produtos = append(produtos, produto)
	}

	return produtos, nil
}
