// Package config provides basic config type
package config

// ServerConfig ...
type ServerConfig struct {
	Host string `yaml:"host"`
	Port int64  `yaml:"port"`

	DB    *DbConfig    `yaml:"db"`
	Redis *RedisConfig `yaml:"redis"`
}
