package log

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v2"
)

var Logger *zap.Logger

type Config struct {
	Logging struct {
		Level            string   `yaml:"level"`
		Encoding         string   `yaml:"encoding"`
		OutputPaths      []string `yaml:"outputPaths"`
		ErrorOutputPaths []string `yaml:"errorOutputPaths"`
		EncoderConfig    struct {
			TimeKey         string `yaml:"timeKey"`
			LevelKey        string `yaml:"levelKey"`
			NameKey         string `yaml:"nameKey"`
			CallerKey       string `yaml:"callerKey"`
			MessageKey      string `yaml:"messageKey"`
			StacktraceKey   string `yaml:"stacktraceKey"`
			LineEnding      string `yaml:"lineEnding"`
			LevelEncoder    string `yaml:"levelEncoder"`
			TimeEncoder     string `yaml:"timeEncoder"`
			DurationEncoder string `yaml:"durationEncoder"`
			CallerEncoder   string `yaml:"callerEncoder"`
		} `yaml:"encoderConfig"`
	} `yaml:"logging"`
}

func InitLogs() {
	var cfg Config

	// Leitura do arquivo de configuração
	configFile, err := os.ReadFile("data/config_log.yaml")
	if err != nil {
		fmt.Printf("Erro ao ler o arquivo de configuração: %v\n", err)
		os.Exit(1)
	}

	// Decodificação do arquivo YAML
	if err := yaml.Unmarshal(configFile, &cfg); err != nil {
		fmt.Printf("Erro ao decodificar o arquivo de configuração: %v\n", err)
		os.Exit(1)
	}

	// Configuração do logger usando as informações do arquivo
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

	Logger, err = config.Build(zap.AddCaller())
	if err != nil {
		panic(err)
	}
}

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
