package types

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Calendar contains semesters, exams and breaks in an academic year
type Calendar struct {
	Semesters Semester `yaml:"semesters"`
	Exams     Exam     `yaml:"exams"`
	Breaks    Break    `yaml:"breaks"`
}

// StartEnd contains when a period starts and ends
type StartEnd struct {
	Start string `yaml:"start"`
	End   string `yaml:"end"`
}

// Semester there are two semesters, winter and sprint
type Semester struct {
	Winter StartEnd `yaml:"winter"`
	Sprint StartEnd `yaml:"sprint"`
}

// Exam there are 3 exam periods, winter, sprint and September
type Exam struct {
	September StartEnd `yaml:"september"`
	Winter    StartEnd `yaml:"winter"`
	Sprint    StartEnd `yaml:"sprint"`
}

// Break contains the periods of the breaks for semester breaks and national holidays
type Break struct {
	Winter  StartEnd  `yaml:"winter"`
	Sprint  StartEnd  `yaml:"sprint"`
	Various []*string `yaml:"various"`
}

// Info contains the booleans to determine the current state of the university
type Info struct {
	IsWeekend      bool `json:"weekend"`
	IsBreak        bool `json:"break"`
	IsExamsPeriod  bool `json:"exams"`
	IsNormalPeriod bool `json:"normal"`
}

// LoadCalendar reads a yaml file and returns calendar struct
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
