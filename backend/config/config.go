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
	HostIp      string `yaml:"hostIp" json:"hostIp" env:"HOST_IP"`
	SitePort    string `yaml:"sitePort" json:"sitePort" env:"SITE_PORT"`
	BackendPort string `yaml:"backendPort" json:"backendPort" env:"BACKEND_PORT"`
}

type HtmlConfig struct {
	Dir  string `yaml:"dir" json:"dir" env:"STATIC_DIR"`
	Load bool   `yaml:"load" json:"load" env:"STATIC_LOAD"`
}

type DBConfig struct {
	Root    string        `yaml:"root" json:"root" env:"DB_ROOT"`
	Timeout time.Duration `yaml:"timeout" json:"timeout" env:"DB_TIMEOUT"`
}

type config struct {
	Version  string     `yaml:"version" json:"version"`
	Location string     `yaml:"location" json:"location" env:"LOCATION"`
	Html     HtmlConfig `yaml:"html" json:"html"`
	Node     NodeConfig `yaml:"node" json:"node"`
	DB       DBConfig   `yaml:"db" json:"db"`
}

func InitConfig() error {
	configuration = new(config)
	if err := readConfigFile("./config/config.yaml"); err != nil {
		return err
	}

	if location := getEnvLocation(); location != "" {
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

func Location() string {
	return configuration.Location
}

func HtmlArgs() HtmlConfig {
	return configuration.Html
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
