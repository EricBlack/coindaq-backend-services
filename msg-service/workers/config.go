package workers

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type ConnInfo struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type ConnList struct {
	List []ConnInfo `yaml:"ConnectionInfo"`
}

type WorkerInfo struct {
	Enable  bool                   `yaml:"enable"`
	Name    string                 `yaml:"name"`
	Queue   string                 `yaml:"queue"`
	Keys    []string               `yaml:"keys"`
	Payload map[string]interface{} `yaml:"payload"`
}

type WorkerList struct {
	List []WorkerInfo `yaml:"WorkerInfo"`
}

func ParseConnInfo() []ConnInfo {
	content, _ := ioutil.ReadFile("workers/config.yaml")
	connList := ConnList{}
	err := yaml.Unmarshal(content, &connList)
	if err != nil {
		fmt.Errorf("Read config got error: %s", err)
	}

	return connList.List

}

func ParseWorkerInfo() []WorkerInfo {
	content, _ := ioutil.ReadFile("workers/config.yaml")
	workerList := WorkerList{}
	err := yaml.Unmarshal(content, &workerList)
	if err != nil {
		fmt.Errorf("Read config got error: %s", err)
	}

	return workerList.List
}
