package types

type Holiday struct {
	Date string `json:"date"`
	Name string `json:"name"`
}

type HolidAPI struct {
	Success  bool       `json:"success"`
	Holidays []*Holiday `json:"holidays"`
}
