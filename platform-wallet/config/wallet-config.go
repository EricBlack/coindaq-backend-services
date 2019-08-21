package config

import (
	"errors"
	"io/ioutil"
	"sync"

	yaml "gopkg.in/yaml.v2"
)

var config map[string](map[string]interface{})
var once sync.Once

func Parse(fpath string) map[string](map[string]interface{}) {
	once.Do(func() {
		config = make(map[string](map[string]interface{}))
		b, err := ioutil.ReadFile(fpath)
		if err != nil {
			panic(err)
		}
		yaml.Unmarshal(b, &config)
	})
	return config
}

func Get() map[string](map[string]interface{}) {
	if len(config) == 0 {
		panic(errors.New("parse config first!"))
	}
	return config
}
