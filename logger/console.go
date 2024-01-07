package logger

import (
	"fmt"
	"sync"
)

var (
	consoleLoggerSingleton iLogger
	consoleLoggerOnce      sync.Once
)

type consoleLogger struct{}

func getConsoleLogger() iLogger {
	consoleLoggerOnce.Do(func() {
		initConsoleLogger()
	})
	return consoleLoggerSingleton
}

func initConsoleLogger() {
	consoleLoggerSingleton = &consoleLogger{}
}

func logMessage(message string, keyValues ...interface{}) {
	if len(keyValues)%2 != 0 {
		fmt.Println("Error: Key-value pairs must be provided in pairs.")
		return
	}

	logMessage := message + " "
	for i := 0; i < len(keyValues); i += 2 {
		key := fmt.Sprintf("%v", keyValues[i])
		value := fmt.Sprintf("%v", keyValues[i+1])
		logMessage += fmt.Sprintf("%s:%s ", key, value)
	}

	fmt.Println(logMessage)
}

func (c *consoleLogger) Debug(message string, keyValues ...interface{}) {
	logMessage(message, keyValues...)
}

func (c *consoleLogger) Info(message string, keyValues ...interface{}) {
	logMessage(message, keyValues...)
}

func (c *consoleLogger) Warn(message string, keyValues ...interface{}) {
	logMessage(message, keyValues...)
}

func (c *consoleLogger) Error(message string, keyValues ...interface{}) {
	logMessage(message, keyValues...)
}
