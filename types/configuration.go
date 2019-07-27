package types

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Configuration contains all the configuration parameters.
type Configuration struct {
	Plugin string `json:"plugin"`
}

type RoomMap struct {
	Rooms map[string]string `yaml:"rooms,omitempty"`
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

func LoadMapping(file string) (*RoomMap, error) {
	byt, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	rooms := &RoomMap{}
	if err := yaml.Unmarshal(byt, rooms); err != nil {
		return nil, err
	}
	return rooms, nil
}
