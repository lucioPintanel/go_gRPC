package configs

import (
	"database/sql"

	logs "grpccli_srv/internal/logs"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./data/database.db")
	if err != nil {
		logs.Logger.Sugar().Errorf("falha em abrir a conexão")
	}

	if err = DB.Ping(); err != nil {
		logs.Logger.Sugar().Errorf("falha em estabelecer a conexão")
	}

	logs.Logger.Debug("Database connection established")
}
