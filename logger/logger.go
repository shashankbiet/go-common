package logger

var (
	defaultLogger   = new()
	defaultLogLevel = LogLevelInfo
)

type appLogger struct {
	logger   iLogger
	logLevel LogLevel
}

func new() *appLogger {
	return &appLogger{
		logger:   getConsoleLogger(),
		logLevel: defaultLogLevel,
	}
}

func InitDefaultLogger(logType LogType, logLevel LogLevel) {
	logger := getLogger(logType)
	defaultLogger = &appLogger{
		logger:   logger,
		logLevel: logLevel,
	}
}

func Debug(message string, keyValues ...interface{}) {
	if defaultLogger.logLevel <= LogLevelDebug {
		defaultLogger.logger.Debug(message, keyValues...)
	}
}

func Info(message string, keyValues ...interface{}) {
	if defaultLogger.logLevel <= LogLevelInfo {
		defaultLogger.logger.Info(message, keyValues...)
	}
}

func Warn(message string, keyValues ...interface{}) {
	if defaultLogger.logLevel <= LogLevelWarn {
		defaultLogger.logger.Warn(message, keyValues...)
	}
}

func Error(message string, keyValues ...interface{}) {
	if defaultLogger.logLevel <= LogLevelError {
		defaultLogger.logger.Error(message, keyValues...)
	}
}

func getLogger(logType LogType) iLogger {
	switch logType {
	case LogTypeZap:
		return getZapLogger()
	case LogTypeConsole:
		return getConsoleLogger()
	default:
		return getConsoleLogger()
	}
}
