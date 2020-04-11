package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Load returns Configuration struct
func Load(path string) (*Configuration, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}
	var cfg = new(Configuration)
	if err := yaml.Unmarshal(bytes, cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}
	return cfg, nil
}

// Configuration holds data for configuring application
type Configuration struct {
	Server *Server   `yaml:"server,omitempty"`
	DB     *Database `yaml:"database,omitempty"`
	URLServices *URLServices `yaml:"url_services"`
}

// Server holds data for server configuration
type Server struct {
	Port         string `yaml:"port,omitempty"`
	Debug        bool   `yaml:"debug,omitempty"`
	ReadTimeout  int    `yaml:"read_timeout_seconds,omitempty"`
	WriteTimeout int    `yaml:"write_timeout_seconds,omitempty"`
}

// Database holds data for database configuration
type Database struct {
	DSN        string `yaml:"dsn,omitempty"`
	LogQueries bool   `yaml:"log_queries,omitempty"`
	Timeout    int    `yaml:"timeout_seconds,omitempty"`
}

//URLServices url address other services
type URLServices struct {
	ClientsSvcBaseURL string `yaml:"client_svc_base_url"`
}
