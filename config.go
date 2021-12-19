// Copyright 2021 AreSZerA. All rights reserved.
// This file initializes config from config.yml and provides functions for getting config values.

package crimson

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// The following structures restrict the YAML configuration file structure.
type config struct {
	Server  configServer   `yaml:"server"`
	Browser configBrowser  `yaml:"browser"`
	Session configSession  `yaml:"session"`
	DB      configDatabase `yaml:"db"`
}

type configServer struct {
	Name    string `yaml:"name"`
	Port    int    `yaml:"port"`
	Timeout int64  `yaml:"timeout"`
}

type configBrowser struct {
	Open bool   `yaml:"open"`
	Page string `yaml:"page"`
}

type configDatabase struct {
	MySQL configMySQL `yaml:"mysql"`
}

type configMySQL struct {
	Driver   string `yaml:"driver"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	URL      string `yaml:"url"`
	Name     string `yaml:"name"`
	Extra    string `yaml:"extra"`
}

type configSession struct {
	Provider string `yaml:"provider"`
	Cookie   string `yaml:"cookie"`
	Timeout  int64  `yaml:"timeout"`
}

// Initialize configuration with default values
var conf = &config{
	Server: configServer{
		Name:    "Crimson",
		Port:    8080,
		Timeout: 60,
	},
	Browser: configBrowser{
		Open: false,
		Page: "/",
	},
	Session: configSession{
		Provider: "Crimson",
		Cookie:   "SESSION_ID",
		Timeout:  3600,
	},
	// No default value for field db
}

func init() {
	// Always read crimson-config.yml
	file, err := ioutil.ReadFile("crimson-config.yml")
	if err != nil {
		PrintWarning("Failed to load config file: " + err.Error())
		return
	}
	// Parse YAML file and reflect to config
	err = yaml.Unmarshal(file, conf)
	if err != nil {
		PrintWarning("Failed to unmarshal config file: " + err.Error())
		return
	}
}

// GetServerName returns application name that configured in config.yml or default value
func GetServerName() string {
	return conf.Server.Name
}

// GetServerPort returns port number that configured in config.yml or default value.
func GetServerPort() int {
	return conf.Server.Port
}

// GetServerTimeout returns server shut down timeout in uint of second that configured in config.yml or default value.
func GetServerTimeout() int64 {
	return conf.Server.Timeout
}

// IsBrowserAutoOpen returns if open browser automatically after starting server
func IsBrowserAutoOpen() bool {
	return conf.Browser.Open
}

// GetBrowserOpenPage returns open page name for browser auto open
func GetBrowserOpenPage() string {
	return conf.Browser.Page
}

// GetMySQLDriver returns MySQL driver name (usually mysql, mariadb is also available for mariadb)
func GetMySQLDriver() string {
	return conf.DB.MySQL.Username
}

// GetMySQLUsername returns MySQL username
func GetMySQLUsername() string {
	return conf.DB.MySQL.Username
}

// GetMySQLPassword returns MySQL password
func GetMySQLPassword() string {
	return conf.DB.MySQL.Password
}

// GetMySQLUrl returns MySQL URL in format of ip_address:port_number
func GetMySQLUrl() string {
	return conf.DB.MySQL.URL
}

func GetMySQLDBName() string {
	return conf.DB.MySQL.Name
}

// GetMySQLExtra returns MySQL extra configs
func GetMySQLExtra() string {
	return conf.DB.MySQL.Extra
}

// GetMySQLDataSrc returns MySQL data source in format of username:password@url/db?extra
func GetMySQLDataSrc() string {
	body := GetMySQLUsername() + ":" + GetMySQLPassword() + "@tcp(" + GetMySQLUrl() + ")/" + GetMySQLDBName()
	if GetMySQLExtra() != "" {
		return body + "?" + GetMySQLExtra()
	}
	return body
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
