package config

import (
	"strings"

	"github.com/spf13/viper"
	"github.com/0had0/5G-core/pkg/common/logger"
	"go.uber.org/zap"
)

// Config holds the configuration for the application
type Config struct {
	// Server configuration
	Server struct {
		Host string
		Port int
		TLS  struct {
			Enabled bool
			Cert    string
			Key     string
		}
	}

	// Database configuration
	Database struct {
		Type     string // "mongodb", "redis", etc.
		Host     string
		Port     int
		Username string
		Password string
		Name     string
	}

	// NRF configuration
	NRF struct {
		URL               string
		RegistrationRetry int
		HeartbeatInterval int
	}

	// Network Function specific configuration
	NetworkFunction struct {
		Type         string // "AMF", "SMF", etc.
		InstanceID   string
		InstanceName string
		Capacity     int
		Priority     int
		Locality     string
	}

	// Kubernetes configuration
	Kubernetes struct {
		Namespace     string
		LabelSelector string
	}

	// Logging configuration
	Logging struct {
		Level string
	}

	// Metrics configuration
	Metrics struct {
		Enabled bool
		Port    int
	}
}

// LoadConfig loads the configuration from environment variables and config files
func LoadConfig(configPath string) (*Config, error) {
	config := &Config{}

	// Set up viper to read from config files and environment variables
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(configPath)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Set default values
	setDefaults(v)

	// Read the config file
	err := v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			logger.Error("Failed to read config file", zap.Error(err))
			return nil, err
		}
		logger.Warn("No config file found, using environment variables and defaults")
	}

	// Unmarshal the configuration
	err = v.Unmarshal(config)
	if err != nil {
		logger.Error("Failed to unmarshal config", zap.Error(err))
		return nil, err
	}

	// Initialize logger with configured log level
	logger.Initialize(config.Logging.Level)

	return config, nil
}

// setDefaults sets default values for the configuration
func setDefaults(v *viper.Viper) {
	// Server defaults
	v.SetDefault("server.host", "0.0.0.0")
	v.SetDefault("server.port", 8080)
	v.SetDefault("server.tls.enabled", false)

	// Database defaults
	v.SetDefault("database.type", "mongodb")
	v.SetDefault("database.host", "localhost")
	v.SetDefault("database.port", 27017)
	v.SetDefault("database.name", "5gcore")

	// NRF defaults
	v.SetDefault("nrf.url", "http://nrf-service:8080")
	v.SetDefault("nrf.registrationRetry", 5)
	v.SetDefault("nrf.heartbeatInterval", 30)

	// Network Function defaults
	v.SetDefault("networkFunction.capacity", 100)
	v.SetDefault("networkFunction.priority", 1)

	// Kubernetes defaults
	v.SetDefault("kubernetes.namespace", "5g-core")

	// Logging defaults
	v.SetDefault("logging.level", "info")

	// Metrics defaults
	v.SetDefault("metrics.enabled", true)
	v.SetDefault("metrics.port", 9090)
}
