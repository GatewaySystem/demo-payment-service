package config

import "os"

type Config struct {
	Port        string
	ServiceName string
	Version     string
}

func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}
	return &Config{
		Port:        port,
		ServiceName: "demo-payment-service",
		Version:     "1.0.0",
	}
}
