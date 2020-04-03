package model

import "time"

//Employment ...
type Employment struct {
	Client *Client
	Date   time.Time
}
