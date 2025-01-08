package configs

import (
	"database/sql"

	logs "grpccli_srv/internal/logs"

	_ "github.com/mattn/go-sqlite3"
)

// DB é a instância global da conexão com o banco de dados.
var DB *sql.DB

// InitDB inicializa a conexão com o banco de dados SQLite.
func InitDB() {
	var err error

	// Abre uma nova conexão com o banco de dados SQLite.
	DB, err = sql.Open("sqlite3", "./data/database.db")
	if err != nil {
		// Log de erro caso falhe ao abrir a conexão.
		logs.Logger.Sugar().Errorf("falha em abrir a conexão")
	}

	// Verifica se a conexão com o banco de dados está ativa.
	if err = DB.Ping(); err != nil {
		// Log de erro caso falhe ao estabelecer a conexão.
		logs.Logger.Sugar().Errorf("falha em estabelecer a conexão")
	}

	// Log de sucesso indicando que a conexão foi estabelecida.
	logs.Logger.Debug("Database connection established")

	// Cria a tabela de produtos no banco de dados.
	createProductTable()
}
