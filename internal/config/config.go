package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/boldlogic/cbr-market-data-worker/pkg/logger"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Log    logger.Config `yaml:"log" json:"log"`
	Server ServerConfig  `yaml:"server" json:"server"`
}

type ServerConfig struct {
	ListenHost   string `yaml:"listen_host" json:"listen_host"`
	ExternalHost string `yaml:"external_host" json:"external_host"`
	Port         int    `yaml:"port" json:"port"`
	Timeout      int    `yaml:"timeout" json:"timeout"`
}

const defaultConfigPath = "config.yaml"

var err error

func ParseConfig() (*Config, error) {
	configPath := flag.String("config", defaultConfigPath, "")
	flag.Parse()

	fileBody, err := os.ReadFile(*configPath)
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать файл конфигурации: %w", err)
	}

	var cfg Config
	if err = yaml.Unmarshal(fileBody, &cfg); err != nil {
		return nil, fmt.Errorf("не удалось разобрать конфигурацию: %w", err)
	}

	return &cfg, nil
}
