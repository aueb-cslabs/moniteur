package types

// Holiday struct that defines a holiday
type Holiday struct {
	Date    string `json:"date"`
	Name    string `json:"localName"`
	IntName string `json:"name"`
}
