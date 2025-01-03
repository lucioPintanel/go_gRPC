package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func init() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.CallerKey = "caller" // Nome do campo para a informação do chamador

	var err error
	Logger, err = config.Build(zap.AddCaller())
	if err != nil {
		panic(err)
	}
}
