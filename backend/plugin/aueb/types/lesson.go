package types

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
