package models

import "time"

type Strength struct {
	Load      string    `json:"Load"`
	CreatedAt time.Time `json:"CreatedAt"`
}

type StrengthAdd struct {
	Id int `json:"Id"`
	Strength
}
