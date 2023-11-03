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
	OrderStatus struct {
		Brokers []string `yaml:"brokers"`
		Topics  []string `yaml:"topics"`
	} `yaml:"order_status"`
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
	c.OrderStatus.Brokers = []string{"kafka-1:9091", "kafka-2:9092", "kafka-3:9093"}
	c.OrderStatus.Topics = []string{"order_status"}
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
