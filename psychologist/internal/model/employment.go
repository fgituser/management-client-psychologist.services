package model

import "time"

//Employment ...
type Employment struct {
	Client  *Client    `json:"client,omitempty"`
	Shedule []*Shedule `json:"shedule,omitempty"`
}

//Shedule ...
type Shedule struct {
	Employee *Employee    `json:"employee,omitempty"`
	DateTime time.Time `json:"date_time,omitempty"`
}
