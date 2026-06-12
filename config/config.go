package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

var Configs *configs

type configs struct {
	ES struct {
		Hosts    []string `yaml:"hosts"`
		User     string   `yaml:"user"`
		Password string   `yaml:"password"`
	} `yaml:"es"`
	Metrics []struct {
		Name     string   `yaml:"name"`
		Index    string   `yaml:"index"`
		Keywords []string `yaml:"keywords"`
		Labels   map[string]string `yaml:"labels"`
	} `yaml:"metrics"`
}

func init() {
	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		panic(fmt.Errorf("read config file failed: %w", err))
	}
	Configs = &configs{}
	if err := yaml.Unmarshal(data, Configs); err != nil {
		panic(err)
	}
}
