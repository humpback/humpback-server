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
	HostIp   string `yaml:"hostIp" json:"hostIp" env:"HOST_IP"`
	SitePort string `yaml:"sitePort" json:"sitePort" env:"SITE_PORT"`
}

type BackendConfig struct {
	BackendPort          string `yaml:"backendPort" json:"backendPort" env:"BACKEND_PORT"`
	CheckInterval        int    `yaml:"checkInterval" json:"checkInterval" env:"BACKEND_CHECKINTERVAL"`
	ServiceCheckInterval int    `yaml:"serviceCheckInterval" json:"serviceCheckInterval" env:"BACKEND_SERVICECHECKINTERVAL"`
	CheckThreshold       int    `yaml:"checkThreshold" json:"checkThreshold" env:"BACKEND_CHECKTHRESHOLD"`
}

type HtmlConfig struct {
	Dir  string `yaml:"dir" json:"dir" env:"STATIC_DIR"`
	Load bool   `yaml:"load" json:"load" env:"STATIC_LOAD"`
}

type DBConfig struct {
	Root    string        `yaml:"root" json:"root" env:"DB_ROOT"`
	Timeout time.Duration `yaml:"timeout" json:"timeout" env:"DB_TIMEOUT"`
}

type AdminConfig struct {
	Id       string `yaml:"id" json:"id" env:"ADMIN_ID"`
	Name     string `yaml:"name" json:"name" env:"ADMIN_NAME"`
	Password string `yaml:"password" json:"password" env:"ADMIN_PASSWORD"`
}

type config struct {
	Version  string        `yaml:"version" json:"version"`
	Location string        `yaml:"location" json:"location" env:"LOCATION"`
	Html     HtmlConfig    `yaml:"html" json:"html"`
	Node     NodeConfig    `yaml:"node" json:"node"`
	DB       DBConfig      `yaml:"db" json:"db"`
	Backend  BackendConfig `yaml:"backend" json:"backend"`
	Admin    AdminConfig   `yaml:"admin" json:"admin"`
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
	//todo 后续添加配置检查，检查不通过则退出程序
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

func BackendArgs() BackendConfig {
	return configuration.Backend
}

func DBArgs() DBConfig {
	return configuration.DB
}

func AdminArgs() AdminConfig {
	return configuration.Admin
}

func Config() any {
	return *configuration
}
