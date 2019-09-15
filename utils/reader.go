package utils

import (
	"github.com/aueb-cslabs/moniteur/types"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Read function that reads existing.yml that contains all the announcements from previous moniteur execution
func Read() (*types.Reader, error) {
	byt, err := ioutil.ReadFile("existing.yml")
	if err != nil {
		return nil, err
	}
	existing := &types.Reader{}
	if err := yaml.Unmarshal(byt, existing); err != nil {
		return nil, err
	}
	return existing, nil
}
