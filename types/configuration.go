package types

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Configuration contains all the configuration parameters.
type Configuration struct {
	Plugin    string `yaml:"plugin"`
	ExamsLink string `yaml:"exams"`
	Secret    string `yaml:"secret"`
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
