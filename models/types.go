package models

type Event struct {
	Id        int16  `json:"eventId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	CDate     string `json:"cdate"`
	Begin     string `json:"begin"`
	End       string `json:"end"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}
