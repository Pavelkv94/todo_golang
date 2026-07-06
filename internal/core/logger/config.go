package logger

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type LoggerConfig struct {
	LogFolder string `envconfig:"FOLDER" required:"true"`
	LogLevel  string `envconfig:"LEVEL" required:"true"`
}

// конструктор для инициализации конфигурации из переменных окружения
func NewConfig() (*LoggerConfig, error) {
	var config LoggerConfig
	if err := envconfig.Process("LOG", &config); err != nil {
		return &LoggerConfig{}, fmt.Errorf("failed to process environment variables: %w", err)
	}
	return &config, nil
}

// конструктор для проверки переменных до запуска приложения
func NewCofigMust() *LoggerConfig {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("failed to load logger config: %w", err)
		panic(err)
	}
	return config
}
