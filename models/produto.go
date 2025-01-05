package models

import "time"

type Produto struct {
	ID          int       `json:"id"`
	Descricao   string    `json:"descricao"`
	DataCriacao time.Time `json:"data_criacao"`
	Categoria   string    `json:"categoria"`
}
