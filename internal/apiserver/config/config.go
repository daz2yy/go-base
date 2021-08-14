package config

import "github.com/daz2yy/go-base/internal/apiserver/options"

// Config is the running configuration structure of the service.
type Config struct {
	*options.Options
}

// CreateConfigFromOptions creates a running configuration instance based
// on a given command line or configuration file option.
func CreateConfigFromOptions(opts *options.Options) (*Config, error) {
	return &Config{opts}, nil
}
