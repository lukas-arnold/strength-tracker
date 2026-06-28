package models

type StrengthInput struct {
	Date        string  `json:"date"`
	Load        float64 `json:"load"`
	Repetitions int64   `json:"repetitions"`
}

type Strength struct {
	Id int64 `json:"id"`
	StrengthInput
}
