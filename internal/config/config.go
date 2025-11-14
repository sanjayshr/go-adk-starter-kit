// Package config provides configuration management for the application
package config

import (
	"flag"
	"log/slog"
	"os"
	"strings"
)

// Config holds application configuration
type Config struct {
	LogLevel    string
	AgentLogger bool
	Prompt      string
}

// New creates a new Config instance with default values
func New() *Config {
	return &Config{
		LogLevel:    "info",
		AgentLogger: true,
		Prompt:      "",
	}
}

// ParseFlags parses command-line flags and returns Config
func ParseFlags() *Config {
	cfg := &Config{}

	flag.StringVar(&cfg.LogLevel, "log-level", "info", "Log level: debug, info, error")
	flag.BoolVar(&cfg.AgentLogger, "agent-logger", true, "Enable agent logger for session state outputs")
	flag.StringVar(&cfg.Prompt, "prompt", "", "Blog prompt to process")

	flag.Parse()
	return cfg
}

// GetLogLevel converts string log level to slog.Level
func (c *Config) GetLogLevel() slog.Level {
	switch strings.ToLower(c.LogLevel) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

// GetAPIKey retrieves the Google API key from environment
func GetAPIKey() string {
	return os.Getenv("GOOGLE_API_KEY")
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	// Add validation logic here if needed
	return nil
}

