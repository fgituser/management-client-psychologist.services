package model

//Employment ...
type Employment struct {
	Client  *Client    `json:"client,omitempty"`
	Shedule []*Shedule `json:"shedule,omitempty"`
}

//Shedule ...
type Shedule struct {
	Date string `json:"date,omitempty"`
	Time string `json:"time,omitempty"`
}
