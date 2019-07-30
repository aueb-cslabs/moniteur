package rest

import (
	"encoding/json"
	"github.com/aueb-cslabs/moniteur/types"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func CalendarInfo(ec echo.Context) error {
	c := ec.(*types.Context)

	info := &types.Info{}

	info.IsWeekend = isWeekend()
	info.IsBreak = isBreak(c.Calendar.Breaks)
	info.IsExamsPeriod = isExams(c.Calendar.Exams)
	info.IsNormalPeriod = isNormal(c.Calendar.Semesters)

	return c.JSON(http.StatusOK, info)
}

func currentDate() string {
	now := time.Now()
	dateF := strconv.Itoa(now.Day()) + "/" + strconv.Itoa(int(now.Month())) + "/" + strconv.Itoa(now.Year())

	return dateF
}

func checkDates(start string, end string) bool {
	startA := strings.Split(start, "/")
	endA := strings.Split(end, "/")

	startY, startM, startD := convertDate(startA)
	endY, endM, endD := convertDate(endA)

	now := time.Now()
	startDate := time.Date(startY, startM, startD, 0, 0, 0, 0, now.Location())
	endDate := time.Date(endY, endM, endD, 0, 0, 0, 0, now.Location())

	return now.After(startDate) && now.Before(endDate)
}

func convertDate(date []string) (int, time.Month, int) {
	day, _ := strconv.Atoi(date[0])
	month, _ := strconv.Atoi(date[1])
	year, _ := strconv.Atoi(date[2])

	return year, time.Month(month), day
}

func isNormal(semester types.Semester) bool {
	start := semester.Winter.Start
	end := semester.Winter.End

	res1 := checkDates(start, end)

	start = semester.Sprint.Start
	end = semester.Sprint.End

	res2 := checkDates(start, end)

	return res1 || res2
}

func isExams(exams types.Exam) bool {
	start := exams.Winter.Start
	end := exams.Winter.End

	res1 := checkDates(start, end)

	start = exams.Sprint.Start
	end = exams.Sprint.End

	res2 := checkDates(start, end)

	return res1 || res2
}

func isBreak(breaks types.Break) bool {
	start := breaks.Winter.Start
	end := breaks.Winter.End

	res1 := checkDates(start, end)

	start = breaks.Sprint.Start
	end = breaks.Sprint.End

	res2 := checkDates(start, end)

	res3 := false
	various := breaks.Various
	current := currentDate()

	res4 := saturateHolidAPI()

	for i := range various {
		if current == *various[i] {
			res3 = true
			break
		}
	}

	return res1 || res2 || res3 || res4
}

func isWeekend() bool {
	now := time.Now()
	day := int(now.Weekday())

	if day == 0 || day == 6 {
		return true
	}
	return false
}

func saturateHolidAPI() bool {
	resp, _ := http.Get("http://api.holid.co/v1/GR/Europe/Athens")
	bts, _ := ioutil.ReadAll(resp.Body)
	var holidAPI *types.HolidAPI
	_ = json.Unmarshal(bts, &holidAPI)

	holidays := holidAPI.Holidays

	if holidAPI.Success && len(holidays) != 0 {
		for i := range holidays {
			if strings.Contains(holidays[i].Name, "Solstice") ||
				strings.Contains(holidays[i].Name, "Equinox") ||
				strings.Contains(holidays[i].Name, "Armed") {
				return false
			}
		}
		return true
	}
	return false
}
