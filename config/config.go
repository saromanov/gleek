// Package config contains definition of configuration
// for the gleek
package config

// Config contains config for the gleek
type Config struct {
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
	User     string `yaml:"user"`
	Address  string `yaml:"address"`
}