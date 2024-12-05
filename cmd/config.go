package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type ShortcutConfig struct {
	Station1 string `yaml:"station1"`
	Station2 string `yaml:"station2"`
}

type Config struct {
	Shortcuts map[string]ShortcutConfig `yaml:"shortcuts"`
}

var config Config

func getConfigPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting home directory: %v\n", err)
		os.Exit(1)
	}
	return filepath.Join(homeDir, ".config", "commandtrein", "config.yaml")
}

func initConfig() {
	configPath := getConfigPath()

	// Create config directory if it doesn't exist
	configDir := filepath.Dir(configPath)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating config directory: %v\n", err)
		os.Exit(1)
	}

	// Read config file if it exists
	data, err := os.ReadFile(configPath)
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "Error reading config file: %v\n", err)
			os.Exit(1)
		}
		// Create default config if file doesn't exist
		config = Config{
			Shortcuts: make(map[string]ShortcutConfig),
		}
		saveConfig()
		return
	}

	// Parse YAML
	if err := yaml.Unmarshal(data, &config); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing config file: %v\n", err)
		os.Exit(1)
	}
}

func saveConfig() {
	configPath := getConfigPath()
	data, err := yaml.Marshal(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling config: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing config file: %v\n", err)
		os.Exit(1)
	}
}

// AddShortcut adds or updates a shortcut in the config
func AddShortcut(name, station1, station2 string) error {
	config.Shortcuts[name] = ShortcutConfig{
		Station1: station1,
		Station2: station2,
	}
	saveConfig()
	return nil
}

// GetShortcut retrieves a shortcut from the config
func GetShortcut(name string) (ShortcutConfig, bool) {
	shortcut, exists := config.Shortcuts[name]
	return shortcut, exists
}
