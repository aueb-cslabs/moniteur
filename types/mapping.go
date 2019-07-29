package types

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type RoomMap struct {
	Rooms map[string]string `yaml:"rooms,omitempty"`
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
