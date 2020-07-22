package types

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Configuration contains all the configuration parameters.
type Configuration struct {
	ExamsLink   string `yaml:"exams"`
	Hostname    string `yaml:"hostname"`
	Secret      string `yaml:"secret"`
	Port        int    `yaml:"port"`
	RedisConfig Redis  `yaml:"redis"`
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
