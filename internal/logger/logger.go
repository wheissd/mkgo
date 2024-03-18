package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Get() *zap.Logger {
	ec := zap.NewDevelopmentEncoderConfig()
	encoder := zapcore.NewJSONEncoder(ec)
	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return zap.New(zapcore.NewCore(encoder, zapcore.Lock(logFile), zapcore.DebugLevel))
}
