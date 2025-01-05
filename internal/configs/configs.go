package configs

import (
	logs "grpccli_srv/internal/logs"
	"os"

	"gopkg.in/yaml.v2"
)

type ConfigDB struct {
	Server struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Database struct {
		Driver     string `yaml:"driver"`
		Datasource string `yaml:"datasource"`
	} `yaml:"database"`
}

var CfgDB ConfigDB

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
