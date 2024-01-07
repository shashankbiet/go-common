package logger

type iLogger interface {
	Debug(message string, keyValues ...interface{})
	Info(message string, keyValues ...interface{})
	Warn(message string, keyValues ...interface{})
	Error(message string, keyValues ...interface{})
}
