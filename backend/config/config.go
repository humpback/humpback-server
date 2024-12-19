package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/caarlos0/env/v11"
	"gopkg.in/yaml.v2"
)

var configuration *config

type NodeConfig struct {
	HostIp string `yaml:"hostIp" json:"hostIp" env:"HOST_IP"`
	Port   string `yaml:"port" json:"port" env:"PORT"`
}

type HtmlDirConfig struct {
	Default string `yaml:"default" json:"default" env:"HTML_DEFAULT_DIR"`
}

type DBConfig struct {
	Root    string        `yaml:"root" json:"root" env:"DB_ROOT"`
	Timeout time.Duration `yaml:"timeout" json:"timeout" env:"DB_TIMEOUT"`
}

type config struct {
	Version  string        `yaml:"version" json:"version"`
	Location string        `yaml:"location" json:"location" env:"LOCATION"`
	HtmlDir  HtmlDirConfig `yaml:"htmlDir" json:"htmlDir"`
	Node     NodeConfig    `yaml:"node" json:"node"`
	DB       DBConfig      `yaml:"db" json:"db"`
}

func InitConfig() error {
	configuration = new(config)
	if err := readConfigFile("./config/config.yaml"); err != nil {
		return err
	}
	location := getEnvLocation()
	if location != "" {
		if err := readConfigFile(fmt.Sprintf("./config/config_%s.yaml", strings.ToLower(location))); err != nil {
			return err
		}
	}
	if err := env.Parse(configuration); err != nil {
		return err
	}
	return configuration.check()
}

func (c *config) check() error {
	return nil
}

func readConfigFile(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("read config file(%s), %s", filePath, err)
	}
	if err = yaml.Unmarshal(data, configuration); err != nil {
		return fmt.Errorf("config file(%s) unmarshal, %s", filePath, err)
	}
	return nil
}

func getEnvLocation() string {
	if value := os.Getenv("LOCATION"); strings.TrimSpace(value) != "" {
		return strings.TrimSpace(value)
	}
	return ""
}

func HtmlDir() HtmlDirConfig {
	return configuration.HtmlDir
}

func NodeArgs() NodeConfig {
	return configuration.Node
}

func DBArgs() DBConfig {
	return configuration.DB
}

func Config() any {
	return *configuration
}
