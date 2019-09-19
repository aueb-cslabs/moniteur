// +build linux darwin

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aueb-cslabs/moniteur/types"
	"github.com/tealeg/xlsx"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Lesson struct {
	Semester        string `json:"semester"`
	LessonComments  string `json:"Lesson_comments"`
	Room            string `json:"Room"`
	LessonTitle     string `json:"Lesson_title"`
	Professor       string `json:"professor"`
	Time            string `json:"time"`
	Day             string `json:"Day"`
	DepartmentTitle string `json:"Department_title"`
}

type RoomMap struct {
	Rooms map[string]string `yaml:"rooms,omitempty"`
}

var (
	// PLUGIN The plugin to be read by the moniteur agent.
	PLUGIN = Plugin{}
	// Provides map for english-greek room name translation
	mapping = &RoomMap{}
	// The xlsx file that AUEB provides that contains the exams schedule
	exams []byte
	// The link that corresponds to the exams file
	link string
)

type Plugin struct {
}

// Initialize Method that initializes crucial functions for the plugin
func (Plugin) Initialize(examsLink string) {
	if len(mapping.Rooms) == 0 {
		mapping, _ = loadMapping("config/mapping.yml")
	}
	link = examsLink
	exams, _ = download()
	configureLink()
	exams, _ = download()
	go scheduleExamsDownload()

	ldapConf = LdapConfiguration{}
	if err := loadLdapConfig(&ldapConf); err != nil {
		panic(err)
	}

}

// Schedule Method that returns current schedule from Schedule Master
func (Plugin) Schedule() (*types.Schedule, error) {
	return retriever(), nil
}

// ScheduleRoom Method that returns current schedule and the room that corresponds to it
func (Plugin) ScheduleRoom(room string) (*types.Schedule, error, string) {
	room, changed := checkMapping(room)
	if !changed {
		room = convertChars(room)
	}
	return retriever(), nil, room
}

// ExamsSchedule Method that returns current exams schedule if that exists
func (Plugin) ExamsSchedule() (*types.Schedule, error) {
	schedule := getExamsSchedule()
	var err error
	if len(exams) == 0 {
		err = errors.New("exams schedule not available")
	}
	return schedule, err
}

// ExamsScheduleRoom Method that returns current exams schedule if that exists and the room that corresponds to it
func (Plugin) ExamsScheduleRoom(room string) (*types.Schedule, error, string) {
	room, changed := checkMapping(room)
	if !changed {
		room = convertChars(room)
	}
	schedule := getExamsSchedule()
	var err error
	if len(exams) == 0 {
		err = errors.New("exams schedule not available")
	}
	return schedule, err, room
}

func (Plugin) AuthorizeUser(username string, password string) (bool, error) {
	return authenticateLdap(username, password)
}

// retriever Method that converts Schedule Master json to our json format
func retriever() *types.Schedule {
	resp := &types.Schedule{}
	for _, lesson := range getEntireSchedule() {
		subject := &types.ScheduleSlot{}
		lessonTime := strings.Split(lesson.Time, "-")
		start, _ := strconv.ParseInt(lessonTime[0], 10, 64)
		end, _ := strconv.ParseInt(lessonTime[1], 10, 64)
		subject.Start = start * int64(3600)
		subject.End = end * int64(3600)
		subject.Room = lesson.Room
		subject.Day = determineDay(lesson.Day)
		subject.Title = lesson.LessonTitle
		subject.Host = lesson.Professor
		subject.Semester = lesson.Semester
		resp.Slots = append(resp.Slots, subject)
	}
	return resp
}

// convertChars Method that converts english characters to greek in order to parse Schedule Master Room Name
func convertChars(room string) string {
	re := regexp.MustCompile("[0-9]+")

	if re.MatchString(room) {
		if strings.Contains(room, "a") {
			room = strings.ReplaceAll(room, "a", "Α")
		}
		if strings.Contains(room, "d") {
			room = strings.ReplaceAll(room, "d", "Δ")
		}
		if strings.Contains(room, "t") {
			room = strings.ReplaceAll(room, "t", "Τ")
		}
		return room
	}

	return room
}

// getEntireSchedule Method that retrieves the schedule for Schedule Master API
func getEntireSchedule() []*Lesson {
	resp, _ := http.Get("http://schedule.aueb.gr/mobile/")
	bts, _ := ioutil.ReadAll(resp.Body)
	var slots []*Lesson
	_ = json.Unmarshal(bts, &slots)
	return slots
}

// getExamsSchedule Method that retrieves the exam schedule from AUEB
func getExamsSchedule() *types.Schedule {
	t := time.Now()
	year := t.Year()
	schedule := &types.Schedule{}

	if len(exams) != 0 {
		file, _ := xlsx.OpenBinary(exams)
		rows := file.Sheet[fmt.Sprintf("%d", year-1)].Rows
		var dayName string
		var day int
		var month int

		for i := 0; i < len(rows); i++ {
			var rooms []string
			var semester string
			var lessonName string
			var examiner string

			if strings.Contains(rows[i].Cells[0].Value, "*") {
				continue
			}

			if strings.Contains(rows[i].Cells[0].Value, "ΗΜΕΡΟΜΗΝΙΑ") {
				continue
			}

			if rows[i].Cells[0].Value == "" {
				continue
			}

			if strings.Contains(rows[i].Cells[4].Value, "ΠΡΥΤΑΝΕΙΑ") {
				break
			}

			if !strings.Contains(rows[i].Cells[1].Value, ":") {
				examsDate := strings.Split(rows[i].Cells[0].Value, " ")
				dayName = examsDate[0]
				day, _ = strconv.Atoi(examsDate[1])
				month = determineMonth(examsDate[2])

			} else {

				rooms = strings.Split(rows[i].Cells[0].Value, ", ")
				timestamp := strings.Split(rows[i].Cells[1].Value, "-")
				semester = rows[i].Cells[2].Value
				lessonName = rows[i].Cells[3].Value
				examiner = rows[i].Cells[4].Value
				start := convertTime(timestamp[0])
				end := convertTime(timestamp[1])
				for j := range rooms {
					slot := &types.ScheduleSlot{}
					slot.Room = rooms[j]
					slot.Day = determineDay(dayName)
					slot.Start = start
					slot.End = end
					slot.Title = lessonName
					slot.Host = examiner
					slot.Semester = semester
					slot.DayNum = day
					slot.MonthNum = int(month)
					schedule.Slots = append(schedule.Slots, slot)
				}
			}
		}
	} else {
		return nil
	}
	return schedule
}

// determineDay Method that converts the Day (from the greek language) to an int
func determineDay(day string) int {
	switch day {
	case "Δευτέρα", "ΔΕΥΤΕΡΑ":
		return 1
	case "Τρίτη", "ΤΡΙΤΗ":
		return 2
	case "Τετάρτη", "ΤΕΤΑΡΤΗ":
		return 3
	case "Πέμπτη", "ΠΕΜΠΤΗ":
		return 4
	case "Παρασκευή", "ΠΑΡΑΣΚΕΥΗ":
		return 5
	case "Σάββατο", "ΣΑΒΒΑΤΟ":
		return 6
	}
	return 0
}

// loadMapping Method that loads the room mapping from english to greek names
func loadMapping(file string) (*RoomMap, error) {
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

// checkMapping Method that checks the room map in order to retrieve the right room name
func checkMapping(room string) (string, bool) {
	rooms := mapping.Rooms

	if rooms != nil && rooms[room] != "" {
		return rooms[room], true
	}

	return room, false
}

// determineMonth Method that converts the Month (from the greek language) to an int
func determineMonth(month string) int {
	switch month {
	case "ΙΑΝΟΥΑΡΙΟΥ":
		return 1
	case "ΦΕΒΡΟΥΑΡΙΟΥ":
		return 2
	case "ΜΑΡΤΙΟΥ":
		return 3
	case "ΑΠΡΙΛΙΟΥ":
		return 4
	case "ΜΑΪΟΥ":
		return 5
	case "ΙΟΥΝΙΟΥ":
		return 6
	case "ΙΟΥΛΙΟΥ":
		return 7
	case "ΑΥΓΟΥΣΤΟΥ":
		return 8
	case "ΣΕΠΤΕΜΒΡΙΟΥ":
		return 9
	case "ΟΚΤΩΒΡΙΟΥ":
		return 10
	case "ΝΟΕΜΒΡΙΟΥ":
		return 11
	case "ΔΕΚΕΜΒΡΙΟΥ":
		return 12
	}
	return 0
}

// convertTime Method that converts time to int64
func convertTime(timestamp string) int64 {
	convert := strings.Split(timestamp, ":")
	var hourF float64

	hourInt, _ := strconv.Atoi(convert[0])

	hourF = float64(hourInt)
	if convert[1] == "30" {
		hourF += .5
	}

	return int64(hourF * 3600)
}

// scheduleExamsDownload Method that checks every so often for new exam schedule
func scheduleExamsDownload() {
	for {
		current := determineNow()

		if current >= 50400 && current <= 54000 {
			configureLink()
			ret, same := download()

			if !same {
				exams = ret
			} else {
				time.Sleep(time.Second * 300)
			}
		} else {
			time.Sleep(time.Hour)
		}
	}
}

// download Method that downloads the exam schedule from AUEB
func download() ([]byte, bool) {
	var ret []byte

	resp, _ := http.Get(link)

	if resp.StatusCode == 200 {
		ret, _ = ioutil.ReadAll(resp.Body)
	} else {
		return exams, true
	}

	return ret, false
}

// determineNow Method that determines the current time as a Unix timestamp
func determineNow() int64 {
	now := time.Now()
	year, month, day := now.Date()
	sod := time.Date(year, month, day, 0, 0, 0, 0, now.Location())
	return int64(now.Sub(sod).Seconds())
}

// configureLink Method that constructs link based on the date
func configureLink() {
	t := time.Now()
	date := t.Format("20060102")
	month := t.Month()
	link = fmt.Sprintf("https://aueb.gr/sites/default/files/aueb/%s_Exams_%s.xlsx", month, date)
}
