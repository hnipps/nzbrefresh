package refresh

import (
	"encoding/json"
	"os"

	"github.com/Tensai75/nzbrefresh/internal/arguments"
)

type Config struct {
	providers []Provider
	arguments.Args
}

type Option func(*Config)

func defaultConfig() *Config {
	return &Config{}
}

func WithNZBFile(path string) Option {
	return func(c *Config) {
		c.NZBFile = path
	}
}
func WithCheckOnly(checkOnly bool) Option {
	return func(c *Config) {
		c.CheckOnly = checkOnly
	}
}
func WithProvider(provider string) Option {
	return func(c *Config) {
		c.Provider = provider
	}
}
func WithDebug(shouldDebug bool) Option {
	return func(c *Config) {
		c.Debug = shouldDebug
	}
}
func WithCsv(writeCsv bool) Option {
	return func(c *Config) {
		c.Csv = writeCsv
	}
}

func loadProviderList(path string) ([]Provider, error) {
	if file, err := os.ReadFile(path); err != nil {
		return nil, err
	} else {
		cfg := Config{}
		if err := json.Unmarshal(file, &cfg.providers); err != nil {
			return nil, err
		}
		return cfg.providers, nil
	}
}
