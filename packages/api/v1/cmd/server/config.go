package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type config struct {
	root            string
	port            uint16
	readTimeout     time.Duration
	writeTimeout    time.Duration
	idleTimeout     time.Duration
	shutdownTimeout time.Duration
}

func loadConfig() (config, error) {
	cfg := config{
		root:            os.Getenv("INCIDENT_ROOT"),
		port:            8080,
		readTimeout:     10 * time.Second,
		writeTimeout:    15 * time.Second,
		idleTimeout:     time.Minute,
		shutdownTimeout: 10 * time.Second,
	}
	if cfg.root == "" {
		cfg.root = "."
	}
	if raw := os.Getenv("PORT"); raw != "" {
		port, err := strconv.ParseUint(raw, 10, 16)
		if err != nil {
			return config{}, fmt.Errorf("PORT must be a decimal port between 0 and 65535")
		}
		cfg.port = uint16(port)
	}
	for _, setting := range []struct {
		name  string
		value *time.Duration
	}{
		{"HTTP_READ_TIMEOUT", &cfg.readTimeout},
		{"HTTP_WRITE_TIMEOUT", &cfg.writeTimeout},
		{"HTTP_IDLE_TIMEOUT", &cfg.idleTimeout},
		{"HTTP_SHUTDOWN_TIMEOUT", &cfg.shutdownTimeout},
	} {
		if raw := os.Getenv(setting.name); raw != "" {
			duration, err := time.ParseDuration(raw)
			if err != nil || duration <= 0 {
				return config{}, fmt.Errorf("%s must be a positive duration, such as 10s", setting.name)
			}
			*setting.value = duration
		}
	}
	return cfg, nil
}
