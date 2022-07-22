package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v3"
)

const DefaultConfigFileName = ".testkube-watch.yaml"

// Resource contains resource configuration
type Resource struct {
	Deployment bool `json:"deployment"`
	Services   bool `json:"svc"`
}

// Config struct contains kubewatch configuration
type Config struct {
	// Resources to watch
	Resource Resource `json:"resource"`

	// For watching specific namespace, leave it empty for watching all.
	// this config is ignored when watching namespaces
	Namespace string `json:"namespace,omitempty"`
}

// New creates new config object
func New() (*Config, error) {
	c := &Config{}
	if err := c.Load(); err != nil {
		return c, err
	}

	return c, nil
}

// Load loads configuration from config file
func (c *Config) Load() error {
	file, err := os.Open(getConfigFile())
	if err != nil {
		return err
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	if len(b) != 0 {
		return yaml.Unmarshal(b, c)
	}

	return nil
}

func getConfigFile() string {
	configFile := filepath.Join(configDir(), DefaultConfigFileName)
	if _, err := os.Stat(configFile); err == nil {
		return configFile
	}

	return ""
}

func configDir() string {
	if configDir := os.Getenv("TKW_HOME"); configDir != "" {
		return configDir
	}

	if runtime.GOOS == "windows" {
		home := os.Getenv("USERPROFILE")
		return home
	}

	return os.Getenv("HOME")
}
