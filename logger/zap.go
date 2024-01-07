package logger

import (
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	zapLoggerSingleton iLogger
	zapLoggerOnce      sync.Once
)

type zapLogger struct {
	sugaredLogger *zap.SugaredLogger
}

func getZapLogger() iLogger {
	zapLoggerOnce.Do(func() {
		initZapLogger()
	})
	return zapLoggerSingleton
}

func initZapLogger() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	logger, err := config.Build()
	if err != nil {
		fmt.Printf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()
	sugar := logger.Sugar()
	zapLoggerSingleton = &zapLogger{
		sugaredLogger: sugar,
	}
}

func (z *zapLogger) Debug(message string, keyValues ...interface{}) {
	z.sugaredLogger.Debug(message, keyValues)
}

func (z *zapLogger) Info(message string, keyValues ...interface{}) {
	z.sugaredLogger.Info(message, keyValues)
}

func (z *zapLogger) Warn(message string, keyValues ...interface{}) {
	z.sugaredLogger.Warn(message, keyValues)
}

func (z *zapLogger) Error(message string, keyValues ...interface{}) {
	z.sugaredLogger.Error(message, keyValues)
}
