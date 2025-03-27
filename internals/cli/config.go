package cli

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Config represents the configuration structure for the application.
type Config struct {
	// RetentionDays specifies how many days to retain data.
	RetentionDays int `yaml:"retention_days"`
}

// LoadConfig reads and parses the configuration file.

// Parameters:
// - path: The file path to the YAML config file.

// Returns:
// - *Config: A pointer to the populated Config struct.
// - error: An error if reading or parsing fails.
func LoadConfig(path string) (*Config, error) {
	// Read the contents of the configuration file.
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read configuration file: %w", err)
	}

	var config Config

	// Unmarshal the YAML data into the Config struct.
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse configuration file: %w", err)
	}

	// Return the loaded configuration on success.
	return &config, nil
}
