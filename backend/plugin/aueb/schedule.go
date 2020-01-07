package main

import (
	"encoding/json"
	"github.com/aueb-cslabs/moniteur/backend/types"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

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
		var sheetName string
		for k := range file.Sheet {
			sheetName = k
			break
		}
		rows := file.Sheet[sheetName].Rows
		var dayName string
		var day int
		var month int

		for i := 5; i < len(rows); i++ {
			var rooms []string
			var semester string
			var lessonName string
			var examiner string

			if len(rows[i].Cells) == 0 {
				continue
			}

			if strings.Contains(rows[i].Cells[0].Value, "*") {
				continue
			}

			if strings.Contains(rows[i].Cells[0].Value, "ΗΜΕΡΟΜΗΝΙΑ") {
				continue
			}

			if rows[i].Cells[0].Value == "" || len(rows[i].Cells[0].Value) == 0 {
				continue
			}

			if strings.Contains(rows[i].Cells[0].Value, "ΠΡΥΤΑΝΕΙΑ") ||
				strings.Contains(rows[i].Cells[1].Value, "ΠΡΥΤΑΝΕΙΑ") ||
				strings.Contains(rows[i].Cells[2].Value, "ΠΡΥΤΑΝΕΙΑ") ||
				strings.Contains(rows[i].Cells[3].Value, "ΠΡΥΤΑΝΕΙΑ") ||
				strings.Contains(rows[i].Cells[4].Value, "ΠΡΥΤΑΝΕΙΑ") {
				break
			}

			if !strings.Contains(rows[i].Cells[1].Value, ":") {
				examsDate := strings.Split(rows[i].Cells[0].Value, " ")
				if len(examsDate) == 3 {
					dayName = examsDate[0]
					day, _ = strconv.Atoi(examsDate[1])
					month = determineMonth(examsDate[2])
				} else {
					continue
				}
			} else {

				rooms = strings.Split(rows[i].Cells[0].Value, ", ")
				timestamp := strings.Split(rows[i].Cells[1].Value, "-")
				semester = rows[i].Cells[2].Value
				lessonName = rows[i].Cells[3].Value
				examiner = rows[i].Cells[4].Value
				dayTimestamp := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Now().Location()).Unix()
				start := convertTime(timestamp[0])
				end := convertTime(timestamp[1])
				for j := range rooms {
					slot := &types.ScheduleSlot{}
					slot.Room = rooms[j]
					slot.Day = determineDay(dayName)
					slot.Start = dayTimestamp + start
					slot.End = dayTimestamp + end
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

// scheduleExamsDownload Method that checks every so often for new exam schedule
func scheduleExamsDownload() {
	for {
		current := determineNow()

		if current >= 50400 && current <= 54000 {
			configureLink()
			ret, same := download()

			if same == nil {
				exams = ret
			} else {
				time.Sleep(time.Second * 300)
			}
		} else {
			time.Sleep(time.Hour)
		}
	}
}

// determineNow Method that determines the current time as a Unix timestamp
func determineNow() int64 {
	now := time.Now()
	year, month, day := now.Date()
	sod := time.Date(year, month, day, 0, 0, 0, 0, now.Location())
	return int64(now.Sub(sod).Seconds())
}
