package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds all configuration for the application
type Config struct {
	RabbitMQ RabbitMQConfig
	MCPServer MCPServerConfig
	Logging   LoggingConfig
	Debug     bool
}

// RabbitMQConfig holds RabbitMQ connection settings
type RabbitMQConfig struct {
	URL      string
	Host     string
	Port     int
	User     string
	Password string
	VHost    string
}

// MCPServerConfig holds MCP server settings
type MCPServerConfig struct {
	Name    string
	Version string
}

// LoggingConfig holds logging settings
type LoggingConfig struct {
	Level  string
	Format string
}

// Load loads configuration from environment variables
func Load() *Config {
	return &Config{
		RabbitMQ: RabbitMQConfig{
			URL:      getEnv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/"),
			Host:     getEnv("RABBITMQ_HOST", "localhost"),
			Port:     getEnvAsInt("RABBITMQ_PORT", 5672),
			User:     getEnv("RABBITMQ_USER", "guest"),
			Password: getEnv("RABBITMQ_PASSWORD", "guest"),
			VHost:    getEnv("RABBITMQ_VHOST", "/"),
		},
		MCPServer: MCPServerConfig{
			Name:    getEnv("MCP_SERVER_NAME", "RabbitMQ MCP Server"),
			Version: getEnv("MCP_SERVER_VERSION", "0.1.0"),
		},
		Logging: LoggingConfig{
			Level:  getEnv("LOG_LEVEL", "info"),
			Format: getEnv("LOG_FORMAT", "json"),
		},
		Debug: getEnvAsBool("DEBUG", false),
	}
}

// GetRabbitMQURL returns the RabbitMQ connection URL
func (c *Config) GetRabbitMQURL() string {
	if c.RabbitMQ.URL != "" {
		return c.RabbitMQ.URL
	}
	return fmt.Sprintf("amqp://%s:%s@%s:%d%s", 
		c.RabbitMQ.User, 
		c.RabbitMQ.Password, 
		c.RabbitMQ.Host, 
		c.RabbitMQ.Port, 
		c.RabbitMQ.VHost)
}

// Helper functions
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
} 