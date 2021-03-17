package models

import "time"

type Logs struct {
	Id int `json:"id"`
	Urls string `json:"urls"`
	Created_at time.Time `json:"created_at"`
}
