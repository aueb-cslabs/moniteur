package utils

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/aueb-cslabs/moniteur/backend/types"
)

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
