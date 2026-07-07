package core_http_server

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type HTTPServerConfig struct {
	Addr            string        `envconfig:"SERVER_ADDR" default:":8080"`
	ShutdownTimeout time.Duration `envconfig:"SERVER_SHUTDOWN_TIMEOUT" default:"10s"`
}

func NewHTTPServerConfig() (HTTPServerConfig, error) {
	var cfg HTTPServerConfig
	if err := envconfig.Process("HTTP", &cfg); err != nil {
		return HTTPServerConfig{}, fmt.Errorf("failed to process HTTP server config: %w", err)
	}
	return cfg, nil
}

func NewConfigMust() HTTPServerConfig {
	cfg, err := NewHTTPServerConfig()
	if err != nil {
		err = fmt.Errorf("failed to get HTTP server config: %w", err)
		panic(err)
	}
	return cfg
}
