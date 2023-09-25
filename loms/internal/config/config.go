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
	Addr string `yaml:"addr"`
}

type Order struct {
	CancelUnpaidTimeout uint64 `yaml:"cancel_unpaid_timeout"`
}

type Config struct {
	Server Server `yaml:"server"`
	Order  Order  `yaml:"order"`
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
	c.Server.Addr = "127.0.0.1:3200"
	c.Order.CancelUnpaidTimeout = 600
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
