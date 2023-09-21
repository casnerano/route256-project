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

type Config struct {
	Server struct {
		Addr string `yaml:"addr"`
	} `yaml:"server"`
	LOMS struct {
		Addr string `yaml:"addr"`
	} `yaml:"LOMS"`
	PIM struct {
		Addr string `yaml:"addr"`
	} `yaml:"PIM"`
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
	c.Server.Addr = ":8080"
	c.LOMS.Addr = "http://loms:8083"
	c.PIM.Addr = "http://route256.pavl.uk:8080"
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
