package models

type ExerciseInput struct {
	Name        string `json:"Name"`
	MuscleGroup string `json:"MuscleGroup"`
}

type Exercise struct {
	Id int64 `json:"Id"`
	ExerciseInput
	StrengthHistory []Strength `json:"StrengthHistory"`
}

type ExerciseWithLastStrength struct {
	Exercise
	LastStrength Strength
}
