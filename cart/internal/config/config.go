// Package config responsible for the app configuration
package config

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const defaultFilename = "./configs/config.yaml"

var yamlFilename string

type Server struct {
	AddrGRPC string `yaml:"addr_grpc"`
	AddrHTTP string `yaml:"addr_http"`
}

type Config struct {
	Server   Server `yaml:"server"`
	Database struct {
		DSN string `yaml:"dsn"`
	} `yaml:"database"`
	LOMS struct {
		Addr string `yaml:"addr"`
	} `yaml:"loms"`
	PIM struct {
		Addr            string `yaml:"addr"`
		RateLimiterAddr string `yaml:"rate_limiter_addr"`
	} `yaml:"pim"`
	Cache struct {
		Shards []string `yaml:"shards"`
	} `yaml:"cache"`
}

func New() (*Config, error) {
	c := &Config{}

	c.SetDefaultValues()

	if err := c.SetFileValues(yamlFilename); err != nil {
		return nil, fmt.Errorf("failed to read configuration file: %w", err)
	}

	return c, nil
}

func (c *Config) SetDefaultValues() {
	c.Server.AddrGRPC = "0.0.0.0:3200"
	c.Server.AddrHTTP = "0.0.0.0:8080"
	c.LOMS.Addr = "loms:3200"
	c.PIM.Addr = "route256.pavl.uk:8082"
	c.PIM.RateLimiterAddr = "pim_rate_limiter:3000"
	c.Cache.Shards = []string{"redis1:6379", "redis2:6379"}
}

func (c *Config) SetFileValues(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(file, c)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	flag.StringVar(&yamlFilename, "config", defaultFilename, "Configuration filename")

	flag.Parse()
}
