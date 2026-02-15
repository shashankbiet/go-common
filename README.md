# go-common

This project aims to create a common library that provides shared utilities for consumption in other Go projects. Currently, it includes a logger module, and more common utilities will be added in the future.

## Features

- Structured logging with key-value pairs
- Multiple log backends: Console and Zap
- Configurable log levels: Debug, Info, Warn, Error

## Tech Stack

- **Language**: Go 1.21.4
- **Logging Library**: [Zap](https://github.com/uber-go/zap) for high-performance structured logging

## Installation

Add the library to your Go project:

```bash
git clone https://github.com/shashankbiet/go-common.git
```

## Environment Variables

This library does not require any environment variables.

## Running the Application

This is a library, not a standalone application. Import it into your Go projects to use the logging functionality.

## Usage Examples

### Logger

```go
package logger

import (
	"sync"

	"github.com/shashankbiet/go-common/logger"
)

var (
	once sync.Once
)

func InitLogger() {
	once.Do(func() {
		logger.InitDefaultLogger(logger.LogTypeZap, logger.LogLevelDebug)
	})
}

func Debug(message string, keyValues ...interface{}) {
	logger.Debug(message, keyValues...)
}

func Info(message string, keyValues ...interface{}) {
	logger.Info(message, keyValues...)
}

func Warn(message string, keyValues ...interface{}) {
	logger.Warn(message, keyValues...)
}

func Error(message string, keyValues ...interface{}) {
	logger.Error(message, keyValues...)
}

```

```go

func InitializeLogger() {
	logger.InitLogger()
	logger.Info("logger setup done")
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.