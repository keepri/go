package config

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/keepri/go/internal/utils/env"
)

type ServerConfig struct {
	Port   int
	Debug  bool
	Logger bool
}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		Port:   env.Get("port", true).Atoi(),
		Debug:  len(env.Get("debug", false).Val) > 0,
		Logger: len(env.Get("logger", false).Val) > 0,
	}
}

func (c *ServerConfig) NewServer(router *chi.Mux) *http.Server {
	return &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%d", c.Port),
	}
}
