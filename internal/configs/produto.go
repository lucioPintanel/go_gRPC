package configs

import logs "grpccli_srv/internal/logs"

// createProductTable cria a tabela de produtos no banco de dados se ela não existir.
func createProductTable() {
	query := `
    CREATE TABLE IF NOT EXISTS produto (
        ID INTEGER PRIMARY KEY AUTOINCREMENT, // ID do produto com autoincremento.
        Descricao TEXT,                       // Descrição do produto.
        DataCriacao DATETIME DEFAULT CURRENT_TIMESTAMP, // Data de criação com valor padrão como timestamp atual.
        Categoria TEXT                        // Categoria do produto.
    );
    `

	// Executa a query para criar a tabela.
	_, err := DB.Exec(query)
	if err != nil {
		// Log de erro caso haja falha na criação da tabela.
		logs.Logger.Sugar().Errorf("Erro ao criar a tabela produto: %v", err)
	}

	// Log de sucesso indicando que a tabela foi verificada/criada.
	logs.Logger.Debug("Tabela produto verificada/criada com sucesso")
}
