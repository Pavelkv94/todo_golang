package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger

	file *os.File
}

func NewLogger(config LoggerConfig) (*Logger, error) {
	zapLvl := zap.NewAtomicLevel()
	if err := zapLvl.UnmarshalText([]byte(config.LogLevel)); err != nil {
		return nil, fmt.Errorf("invalid log level: %w", err)
	}

	if err := os.MkdirAll(config.LogFolder, 0755); err != nil {
		return nil, fmt.Errorf("failed to create log directory: %w", err)
	}

	timestamp := time.Now().UTC().Format("2006-01-02T15-04-05.000000")
	logfilePath := filepath.Join(config.LogFolder, fmt.Sprintf("-%s.log", timestamp))

	file, err := os.OpenFile(logfilePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	zapConfig := zap.NewDevelopmentEncoderConfig()
	zapConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05.000000Z")

	zapEncoder := zapcore.NewConsoleEncoder(zapConfig)

	core := zapcore.NewTee(
		zapcore.NewCore(zapEncoder, zapcore.AddSync(os.Stdout), zapLvl), // вывод в консоль
		zapcore.NewCore(zapEncoder, zapcore.AddSync(file), zapLvl),      // вывод в файл
	)

	logger := zap.New(core, zap.AddCaller())

	return &Logger{
		Logger: logger,
		file:   file,
	}, nil
}

func (l *Logger) Close() error {
	if err := l.file.Close(); err != nil {
		return fmt.Errorf("failed to close log file: %w", err)
	}
	return nil
}
