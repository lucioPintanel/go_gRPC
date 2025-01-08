package log

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v2"
)

// Logger é a instância global do logger.
var Logger *zap.Logger

// Config representa a estrutura de configuração para o logger.
type Config struct {
	Logging struct {
		Level            string   `yaml:"level"`            // Nível de log (debug, info, warn, error).
		Encoding         string   `yaml:"encoding"`         // Formato de saída do log (json, console).
		OutputPaths      []string `yaml:"outputPaths"`      // Caminhos dos arquivos de saída dos logs.
		ErrorOutputPaths []string `yaml:"errorOutputPaths"` // Caminhos dos arquivos de saída dos logs de erro.
		EncoderConfig    struct {
			TimeKey         string `yaml:"timeKey"`         // Chave do tempo.
			LevelKey        string `yaml:"levelKey"`        // Chave do nível de log.
			NameKey         string `yaml:"nameKey"`         // Chave do nome.
			CallerKey       string `yaml:"callerKey"`       // Chave do chamador.
			MessageKey      string `yaml:"messageKey"`      // Chave da mensagem.
			StacktraceKey   string `yaml:"stacktraceKey"`   // Chave do stacktrace.
			LineEnding      string `yaml:"lineEnding"`      // Fim de linha.
			LevelEncoder    string `yaml:"levelEncoder"`    // Codificador de nível.
			TimeEncoder     string `yaml:"timeEncoder"`     // Codificador de tempo.
			DurationEncoder string `yaml:"durationEncoder"` // Codificador de duração.
			CallerEncoder   string `yaml:"callerEncoder"`   // Codificador de chamador.
		} `yaml:"encoderConfig"`
	} `yaml:"logging"`
}

// InitLogs inicializa o logger com base no arquivo de configuração.
func InitLogs() {
	var cfg Config

	// Leitura do arquivo de configuração.
	configFile, err := os.ReadFile("data/config.yaml")
	if err != nil {
		fmt.Printf("Erro ao ler o arquivo de configuração: %v\n", err)
		os.Exit(1)
	}

	// Decodificação do arquivo YAML.
	if err := yaml.Unmarshal(configFile, &cfg); err != nil {
		fmt.Printf("Erro ao decodificar o arquivo de configuração: %v\n", err)
		os.Exit(1)
	}

	// Configuração do logger usando as informações do arquivo.
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(getLogLevel(cfg.Logging.Level))
	config.Encoding = cfg.Logging.Encoding
	config.OutputPaths = cfg.Logging.OutputPaths
	config.ErrorOutputPaths = cfg.Logging.ErrorOutputPaths
	config.EncoderConfig = zapcore.EncoderConfig{
		TimeKey:        cfg.Logging.EncoderConfig.TimeKey,
		LevelKey:       cfg.Logging.EncoderConfig.LevelKey,
		NameKey:        cfg.Logging.EncoderConfig.NameKey,
		CallerKey:      cfg.Logging.EncoderConfig.CallerKey,
		MessageKey:     cfg.Logging.EncoderConfig.MessageKey,
		StacktraceKey:  cfg.Logging.EncoderConfig.StacktraceKey,
		LineEnding:     cfg.Logging.EncoderConfig.LineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout(time.RFC3339),
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// Cria a instância do logger.
	Logger, err = config.Build(zap.AddCaller())
	if err != nil {
		panic(err)
	}
}

// getLogLevel converte o nível de log da configuração para zapcore.Level.
func getLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}
