package models

type StrengthInput struct {
	Date        string  `json:"Date"`
	Load        float64 `json:"Load"`
	Repetitions int64   `json:"Repetitions"`
}

type Strength struct {
	Id int64 `json:"Id"`
	StrengthInput
}
