package logger

type LogType int

const (
	LogTypeZap LogType = iota
	LogTypeConsole
)
