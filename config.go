// Copyright 2021 AreSZerA. All rights reserved.
// This file initializes config from config.yml and provides functions for getting config values.

package crimson

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// config, configServer,and configSession restricts the YAML configuration file structure.
type config struct {
	Server  configServer  `yaml:"server"`
	Session configSession `yaml:"session"`
}

// config, configServer,and configSession restricts the YAML configuration file structure.
type configServer struct {
	Port     int   `yaml:"port"`
	Timeout  int64 `yaml:"timeout"`
	AutoOpen bool  `yaml:"auto-open"`
}

// config, configServer,and configSession restricts the YAML configuration file structure.
type configSession struct {
	Provider string `yaml:"provider"`
	Cookie   string `yaml:"cookie"`
	Timeout  int64  `yaml:"timeout"`
}

// Initialize configuration with default values
var conf = &config{
	Server: configServer{
		Port:     8080,
		Timeout:  60,
		AutoOpen: false,
	},
	Session: configSession{
		Provider: "Crimson",
		Cookie:   "SESSION_ID",
		Timeout:  3600,
	},
}

func init() {
	// Always read config.yml
	file, err := ioutil.ReadFile("config.yml")
	if err != nil {
		PrintError("Failed to load config file: " + err.Error())
		os.Exit(4)
	}
	// Parse YAML file and reflect to config
	err = yaml.Unmarshal(file, conf)
	if err != nil {
		PrintError("Failed to unmarshal config file: " + err.Error())
		os.Exit(4)
	}
}

// GetServerPort returns port number that configured in config.yml or default value.
func GetServerPort() int {
	return conf.Server.Port
}

// GetServerTimeout returns server shut down timeout in uint of second that configured in config.yml or default value.
func GetServerTimeout() int64 {
	return conf.Server.Timeout
}

func IsBrowserAutoOpen() bool {
	return conf.Server.AutoOpen
}

// GetSessionProviderName returns session provider name that configured in config.yml or default value.
func GetSessionProviderName() string {
	return conf.Session.Provider
}

// GetSessionCookieName returns cookie name for session that configured in config.yml or default value.
func GetSessionCookieName() string {
	return conf.Session.Cookie
}

// GetSessionTimeout returns session maximum life time that configured in config.yml or default value.
func GetSessionTimeout() int64 {
	return conf.Session.Timeout
}
