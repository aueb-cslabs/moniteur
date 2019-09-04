package types

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// RoomMap contains a map for english-greek room name mapping
type RoomMap struct {
	Rooms map[string]string `yaml:"rooms,omitempty"`
}

// LoadMapping reads a yaml file and returns RoomsMap struct
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
