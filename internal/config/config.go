package config

import (
	"os"
	"strconv"

	"github.com/keepri/rss/internal/utils"
)

type Config struct {
	Port   int
	Debug  bool
	Logger bool
}

func New() *Config {
	return &Config{
		Port:   8080,
		Debug:  false,
		Logger: false,
	}
}

func (c *Config) LoadFromEnv() *Config {
	port, port_err := strconv.Atoi(os.Getenv("PORT"))
	if port_err == nil {
		c.Port = port
	}

	debug := utils.Env("DEBUG", false)
	if len(debug) > 0 {
		c.Debug = true
	}

	logger := utils.Env("LOGGER", false)
	if len(logger) > 0 {
		c.Logger = true
	}

	return c
}
