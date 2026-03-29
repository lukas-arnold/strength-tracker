package models

import "time"

type StrengthAdd struct {
	Load      string    `json:"Load"`
	CreatedAt time.Time `json:"CreatedAt"`
}

type StrengthGet struct {
	Id int `json:"Id"`
	StrengthAdd
}
