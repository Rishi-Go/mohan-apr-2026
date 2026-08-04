package logger

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func ZapLogger() {

	zapFile, err := os.OpenFile("zap_log_file.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer zapFile.Close()

	enconfig := zap.NewProductionEncoderConfig()

	logger := zap.New(
		zapcore.NewTee(
			zapcore.NewCore(
				zapcore.NewJSONEncoder(enconfig),
				zapcore.Lock(zapFile),
				zapcore.DebugLevel,
			),
			zapcore.NewCore(
				zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
				zapcore.Lock(os.Stdout),
				zapcore.DebugLevel,
			),
		),
	)

	logger.Debug("we in DEBUG mode")
	logger.Info("we in INFO mode")
	logger.Warn("we in WARN mode")
	logger.Error("we in ERROR mode")
}
