package utils

import (
	"github.com/aueb-cslabs/moniteur/types"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Read function that reads existing.yml that contains all the announcements from previous moniteur execution
func Read() (*types.Reader, error) {
	byt, err := ioutil.ReadFile("config/existing.yml")
	if err != nil {
		err = write(&types.Reader{}, "config/existing.yml")
		return nil, err
	}
	existing := &types.Reader{}
	if err := yaml.Unmarshal(byt, existing); err != nil {
		return nil, err
	}
	return existing, nil
}
