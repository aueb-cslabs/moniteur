package types

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Configuration contains all the configuration parameters.
type Configuration struct {
	Plugin      string `yaml:"plugin"`
	ExamsLink   string `yaml:"exams"`
	Hostname    string `yaml:"hostname"`
	Secret      string `yaml:"secret"`
	Port        int    `yaml:"port"`
	RedisConfig Redis  `yaml:"redis"`
	Postgres    string `yaml:"postgres"`
}

// LoadConfiguration reads a configuration file and returns a struct.
func LoadConfiguration(file string) (*Configuration, error) {
	byt, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	config := &Configuration{}
	if err := yaml.Unmarshal(byt, config); err != nil {
		return nil, err
	}
	return config, nil
}
