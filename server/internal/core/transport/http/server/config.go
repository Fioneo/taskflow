package core_http_server

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type ServerConfig struct {
	Addr            string        `envconfig:"ADDR" required:"true"`
	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" required:"true"`
}

func NewConfig() (ServerConfig, error) {
	var config ServerConfig

	if err := envconfig.Process("HTTP", &config); err != nil {
		return ServerConfig{}, fmt.Errorf("process envconfig: %w", err)
	}
	return config, nil
}
func NewConfigMust() ServerConfig {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("get HTTP server config: %w", err)
		panic(err)
	}
	return config
}
