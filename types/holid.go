package types

// Holiday struct that defines a holiday
type Holiday struct {
	Date string `json:"date"`
	Name string `json:"name"`
}

// HolidAPI struct of Holid API saturation
type HolidAPI struct {
	Success  bool       `json:"success"`
	Holidays []*Holiday `json:"holidays"`
}
