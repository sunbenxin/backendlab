// Package config provides basic config type
package config

// RedisConfig ...
type RedisConfig struct {
	// example: redis://host:port
	/*
		redis://HOST[:PORT][?db=DATABASE[&password=PASSWORD]]
		redis://HOST[:PORT][?password=PASSWORD[&db=DATABASE]]
		redis://[:PASSWORD@]HOST[:PORT][/DATABASE]
		redis://[:PASSWORD@]HOST[:PORT][?db=DATABASE]
		redis://HOST[:PORT]/DATABASE[?password=PASSWORD]
	*/
	Host string `yaml:"host"`
}
