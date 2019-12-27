package utils

import (
	"github.com/aueb-cslabs/moniteur/backend/types"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// UpdateUsers updates the authorized users in config.yml
func UpdateUsers(users map[string]string) error {
	config, _ := types.LoadConfiguration("config/config.yml")
	config.AuthorizedUsers = users
	return write(config, "config/config.yml")
}

func UpdateCalendar(calendar *types.Calendar) error {
	return write(calendar, "config/calendar.yml")
}

// write an interface and marshals it as a yml file
func write(msg interface{}, file string) error {
	byt, err := yaml.Marshal(msg)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(file, byt, 0666)
}
