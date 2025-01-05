package configs

import logs "grpccli_srv/internal/logs"

func createProductTable() {
	query := `
    CREATE TABLE IF NOT EXISTS produto (
        ID INTEGER PRIMARY KEY AUTOINCREMENT,
        Descricao TEXT,
        DataCriacao DATETIME DEFAULT CURRENT_TIMESTAMP,
        Categoria TEXT
    );
    `

	_, err := DB.Exec(query)
	if err != nil {
		logs.Logger.Sugar().Errorf("Erro ao criar a tabela produto: %v", err)
	}

	logs.Logger.Debug("Tabela produto verificada/criada com sucesso")
}
