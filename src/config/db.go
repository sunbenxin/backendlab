// Package config provides basic config type
package config

// DbConfig ...
type DbConfig struct {
	Host string `yaml:"host"`
	Port int64  `yaml:"port"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
}
