package types

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Calendar struct {
	Semesters Semester `yaml:"semesters"`
	Exams     Exam     `yaml:"exams"`
	Breaks    Break    `yaml:"breaks"`
}

type StartEnd struct {
	Start string `yaml:"start"`
	End   string `yaml:"end"`
}

type Semester struct {
	Winter StartEnd `yaml:"winter"`
	Sprint StartEnd `yaml:"sprint"`
}

type Exam struct {
	September StartEnd `yaml:"september"`
	Winter    StartEnd `yaml:"winter"`
	Sprint    StartEnd `yaml:"sprint"`
}

type Break struct {
	Winter  StartEnd  `yaml:"winter"`
	Sprint  StartEnd  `yaml:"sprint"`
	Various []*string `yaml:"various"`
}

type Info struct {
	IsWeekend      bool `json:"weekend"`
	IsBreak        bool `json:"break"`
	IsExamsPeriod  bool `json:"exams"`
	IsNormalPeriod bool `json:"normal"`
}

func LoadCalendar(file string) (*Calendar, error) {
	byt, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	calendar := &Calendar{}
	if err := yaml.Unmarshal(byt, calendar); err != nil {
		return nil, err
	}
	return calendar, err
}
