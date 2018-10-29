package server

import (
	"github.com/miracl/conflate"
)

// Config defines the main settings of the game server
type Config struct {
	Address string
}

// LoadConfig loads a config from a local file or url
func LoadConfig(paths ...string) (*Config, error) {
	c, err := conflate.FromFiles(paths...)
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = c.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
