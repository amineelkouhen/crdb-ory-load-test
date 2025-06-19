package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Hydra struct {
		AdminAPI *string `yaml:"admin_api",omitempty`
		PublicAPI  *string `yaml:"public_api",omitempty`
	} `yaml:"hydra"`

	Kratos struct {
		AdminAPI *string `yaml:"admin_api",omitempty`
		PublicAPI  *string `yaml:"public_api",omitempty`
	} `yaml:"kratos"`

	Keto struct {
		WriteAPI *string `yaml:"write_api",omitempty`
		ReadAPI  *string `yaml:"read_api",omitempty`
	} `yaml:"keto"`

	Workload struct {
		ReadRatio       int `yaml:"read_ratio"`
        DurationSec     int `yaml:"duration_sec"`
	} `yaml:"workload"`
}

var AppConfig Config

func LoadConfig(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	if err := yaml.Unmarshal(data, &AppConfig); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return nil
}
