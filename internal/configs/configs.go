package configs

import (
	logs "grpccli_srv/internal/logs"
	"os"

	"gopkg.in/yaml.v2"
)

// ConfigDB define a estrutura de configuração para o banco de dados e servidor.
type ConfigDB struct {
	Server struct {
		Port int    `yaml:"port"` // Porta do servidor.
		Host string `yaml:"host"` // Host do servidor.
	} `yaml:"server"`
	Database struct {
		Driver     string `yaml:"driver"`     // Driver do banco de dados.
		Datasource string `yaml:"datasource"` // Fonte de dados do banco de dados.
	} `yaml:"database"`
}

// CfgDB é a instância global da configuração do banco de dados.
var CfgDB ConfigDB

// LoadConfig carrega a configuração do arquivo data/config.yaml.
func LoadConfig() {
	data, err := os.ReadFile("data/config.yaml")
	if err != nil {
		logs.Logger.Sugar().Errorf("error: %v", err)
	}
	err = yaml.Unmarshal(data, &CfgDB)
	if err != nil {
		logs.Logger.Sugar().Errorf("error: %v", err)
	}
}
