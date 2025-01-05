package repositories

import (
	"grpccli_srv/internal/configs"
	logs "grpccli_srv/internal/logs"
	"grpccli_srv/models"
)

func CreateProduto(produto models.Produto) error {
	query := `
    INSERT INTO produto (Descricao, DataCriacao, Categoria) 
    VALUES (?, ?, ?)
    `
	_, err := configs.DB.Exec(query, produto.Descricao, produto.DataCriacao, produto.Categoria)
	if err != nil {
		logs.Logger.Error("falha em executar a query")
		return err
	}
	return nil
}

func GetProdutoByID(id int) (models.Produto, error) {
	query := "SELECT ID, Descricao, DataCriacao, Categoria FROM produto WHERE ID = ?"
	row := configs.DB.QueryRow(query, id)

	var produto models.Produto
	err := row.Scan(&produto.ID, &produto.Descricao, &produto.DataCriacao, &produto.Categoria)
	if err != nil {
		logs.Logger.Sugar().Errorf("falha em obter o produto - ID: %d", id)
		return produto, err
	}

	return produto, nil
}

func UpdateProduto(produto models.Produto) error {
	query := `
    UPDATE produto 
    SET Descricao = ?, DataCriacao = ?, Categoria = ? 
    WHERE ID = ?
    `
	_, err := configs.DB.Exec(query, produto.Descricao, produto.DataCriacao, produto.Categoria, produto.ID)
	if err != nil {
		logs.Logger.Sugar().Errorf("falha em fazer update - ID: %d", produto.ID)
		return err
	}
	return nil
}

func DeleteProduto(id int) error {
	query := "DELETE FROM produto WHERE ID = ?"
	_, err := configs.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func GetAllProdutos() ([]models.Produto, error) {
	query := "SELECT ID, Descricao, DataCriacao, Categoria FROM produto"
	rows, err := configs.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var produtos []models.Produto
	for rows.Next() {
		var produto models.Produto
		if err := rows.Scan(&produto.ID, &produto.Descricao, &produto.DataCriacao, &produto.Categoria); err != nil {
			return nil, err
		}
		produtos = append(produtos, produto)
	}

	return produtos, nil
}
