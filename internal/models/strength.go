package models

import "time"

type Strength struct {
	CreatedAt time.Time `json:"CreatedAt"`
	Load      string    `json:"Load"`
}
