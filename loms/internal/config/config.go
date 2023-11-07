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

type Order struct {
	CancelUnpaidTimeout uint64 `yaml:"cancel_unpaid_timeout"`
	StatusSender        struct {
		Brokers []string `yaml:"brokers"`
		Topic   string   `yaml:"topic"`
	} `yaml:"status_sender"`
}

type Config struct {
	Server   Server `yaml:"server"`
	Database struct {
		DSN string `yaml:"dsn"`
	} `yaml:"database"`
	Order Order `yaml:"order"`
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
	c.Order.CancelUnpaidTimeout = 30
	c.Order.StatusSender.Brokers = []string{"kafka-1:9091", "kafka-2:9092", "kafka-3:9093"}
	c.Order.StatusSender.Topic = "order_status"
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
