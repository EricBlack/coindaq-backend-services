package config

import (
	"errors"
	"io/ioutil"
	"sync"

	yaml "gopkg.in/yaml.v2"
)

type DataSource struct {
	DriverName string `yaml:"driver_name"`
	URI        string `yaml:"uri"`
	MaxIdle    int    `yaml:"max_idle"`
	MaxOpen    int    `yaml:"max_open"`
	ShowSQL    bool   `yaml:"show_sql"`
}

type EmailSource struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type AppConfig struct {
	Mode        string      `yaml:"mode"`
	Name        string      `yaml:"name"`
	Protocal    string      `yaml:"protocal"`
	Domain      string      `yaml:"domain"`
	CtxPath     string      `yaml:"ctxpath"`
	Addr        string      `yaml:"addr"`
	Port        int         `yaml:"port"`
	Datasource  DataSource  `yaml:"datasource"`
	EmailSource EmailSource `yaml:"emailsource"`
}

var config AppConfig
var once sync.Once

func Parse(fpath string) AppConfig {
	once.Do(func() {
		config = AppConfig{}
		b, err := ioutil.ReadFile(fpath)
		if err != nil {
			panic(err)
		}
		yaml.Unmarshal(b, &config)
	})
	return config
}

func Get() AppConfig {
	if config.Mode == "" {
		panic(errors.New("parse config first!"))
	}
	return config
}
