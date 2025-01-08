package models

import "time"

// Produto representa a estrutura do produto no banco de dados.
type Produto struct {
	ID          int       `json:"id"`           // ID do produto.
	Descricao   string    `json:"descricao"`    // Descrição do produto.
	DataCriacao time.Time `json:"data_criacao"` // Data de criação do produto.
	Categoria   string    `json:"categoria"`    // Categoria do produto.
}
